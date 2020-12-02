package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"sort"
	"strconv"
	"sync"

	"time"

	pb "protobuf/synergy"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/polarbroadband/goto/util"
	log "github.com/sirupsen/logrus"
)

var (
	// local host name
	HOST, _ = os.Hostname()
	// database
	DATA_DB = "mongodb://" + os.Getenv("CRED_UID_DATA") + ":" + os.Getenv("CRED_PWD_DATA") + "@" + os.Getenv("HOST_DATA_DB") + ":27017"
)

func init() {
	// config package level default logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

type Runner struct {
	ID      string
	Ctx     context.Context
	Cancel  context.CancelFunc
	Outlets []*chan pb.RunnerOutput
	oLock   *sync.Mutex
	Wkr     pb.WorkerClient
	Log     *log.Entry
}

func (runner *Runner) Init() error {
	_e := util.NewExeErr("Init", runner.ID)

	result, err := runner.Wkr.CreateRunner(runner.Ctx, &pb.RunnerRequest{ID: runner.ID})
	if err != nil {
		runner.Log.WithError(err).Error("failed worker request, unable to create service")
		return _e.Error("unable to create service on worker", err)
	}
	runner.Log.Infof("worker service created %+v", result)
	wkrStream, err := runner.Wkr.ConnectRunner(runner.Ctx, &pb.RunnerUpdate{ID: runner.ID})
	if err != nil {
		runner.Log.WithError(err).Error("failed worker request, unable to connect service")
		return _e.Error("unable to connect service on worker", err)
	}
	go func() {
		defer runner.Cancel()
		abort := false
		for !abort {
			select {
			case <-runner.Ctx.Done():
				runner.Log.Info("runner closed: cancelled")
				return
			default:
				msg, err := wkrStream.Recv()
				if err != nil {
					abort = true
					if err == io.EOF {
						msg = &pb.RunnerOutput{Msg: "worker stream end"}
						runner.Log.Infof("runner closed: %s", msg.Msg)
					} else {
						msg = &pb.RunnerOutput{Msg: "worker streaming error: " + err.Error()}
						runner.Log.Errorf("runner closed: %s", msg.Msg)
					}
				}
				for ci, ch := range runner.Outlets {
					select {
					case <-*ch:
						runner.Log.Infof("session %v disconnected", ci)
						continue
					default:
						select {
						case *ch <- *msg:
							fmt.Printf("Sent %v to channel %v\n", msg.Msg, ci)
						default:
							fmt.Printf("no receiver on channel %v, Drop %v\n", ci, msg.Msg)
						}
					}
				}
			}
		}
	}()
	runner.Log.Info("runner start")
	return nil
}

type Worker struct {
	*pb.SvrStat
	Conn *grpc.ClientConn
}

type Workers []*Worker

// function to be called by sort package
func (w Workers) Len() int {
	return len(w)
}
func (w Workers) Less(i, j int) bool {
	// sort by Workload, highest first
	return w[i].GetLoad() > w[j].GetLoad()
}
func (w Workers) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type Pool struct {
	Workers
	pLock *sync.Mutex
}

func (p *Pool) Sort() {
	p.pLock.Lock()
	defer p.pLock.Unlock()
	sort.Sort(p.Workers)
}
func (p *Pool) Find(h string) (bool, int) {
	for i, w := range p.Workers {
		if w.GetHost() == h {
			return true, i
		}
	}
	return false, 99999
}

// caller should ensure the worker uniqueness
func (p *Pool) Add(wkr *Worker) {
	p.pLock.Lock()
	defer p.pLock.Unlock()
	p.Workers = append(p.Workers, wkr)
}
func (p *Pool) Delete(h string) {
	if exist, i := p.Find(h); exist {
		// associated services will be closed due to worker gRPC connection close
		p.Workers[i].Conn.Close()
		p.pLock.Lock()
		p.Workers = append(p.Workers[:i], p.Workers[i+1:]...)
		p.pLock.Unlock()
	}
}

func (p *Pool) Sanity() error {
	available := false
	for _, w := range p.Workers {
		wkr := pb.NewWorkerClient(w.Conn)
		gCtx, gCancel := context.WithTimeout(context.Background(), time.Second*10)
		state, err := wkr.Healtz(gCtx, &pb.Request{})
		gCancel()
		if err != nil {
			p.Delete(w.GetHost())
			continue
		}
		p.pLock.Lock()
		w.SvrStat = state
		p.pLock.Unlock()
		if state.GetLoad() <= 85 { // load < 85
			available = true
		}
	}
	p.Sort()
	fmt.Printf("\n--- registrated workers --- %+v\n", p.Workers)
	if !available || len(p.Workers) == 0 {
		return fmt.Errorf("resource depleted")
	}
	return nil
}

type Nest struct {
	Runners map[string]*Runner
	nLock   *sync.Mutex
}

// gc removes the idle runner instance from the nest
func (nest *Nest) gc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for id, runner := range nest.Runners {
				select {
				case <-runner.Ctx.Done():
					nest.nLock.Lock()
					delete(nest.Runners, id)
					nest.nLock.Unlock()
					runner.Log.Warnf("closed, (%v) runners remain active", len(nest.Runners))
				default:
				}
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func (nest *Nest) Add(r *Runner) {
	nest.nLock.Lock()
	defer nest.nLock.Unlock()
	nest.Runners[r.ID] = r
}

type Controller struct {
	pb.UnimplementedControllerServer
	Nest
	Pool
	Log *log.Entry
}

func (c *Controller) Schedule() (pb.WorkerClient, error) {
	if err := c.Pool.Sanity(); err != nil {
		return nil, err
	}
	return pb.NewWorkerClient(c.Pool.Workers[len(c.Pool.Workers)-1].Conn), nil
}

func (c *Controller) CreateRunner(ctx context.Context, r *pb.RunnerRequest) (*pb.SvcStat, error) {
	_e := util.NewExeErr("Create", "runner instance")

	rid := strconv.Itoa(time.Now().Nanosecond())
	if r.GetID() == "" {
		if r.GetTask() == "" {
			return nil, status.Errorf(codes.InvalidArgument, _e.String("invalid request", "missing task config"))
		}
		// create new with provided task config
	} else if _, exist := c.Nest.Runners[r.GetID()]; exist {
		return nil, status.Errorf(codes.InvalidArgument, _e.String("invalid request", "active runner"))
	} else {
		// copy referenced runner task config
	}

	// locate worker resource
	wkr, err := c.Schedule()
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, _e.String("unable to locate worker resource", err))
	}

	// create runner record in database, get runner id

	newRunner := Runner{
		ID:      rid,
		Wkr:     wkr,
		Log:     c.Log.WithField("runner", rid),
		Outlets: []*chan pb.RunnerOutput{},
		oLock:   &sync.Mutex{},
	}
	newRunner.Ctx, newRunner.Cancel = context.WithCancel(context.Background())
	if err := newRunner.Init(); err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	c.Nest.Add(&newRunner)
	c.Log.Infof("new runner created, total: %v", len(c.Nest.Runners))
	return &pb.SvcStat{ID: rid, State: "create"}, nil
}

func (c *Controller) ConnectRunner(r *pb.RunnerUpdate, stream pb.Controller_ConnectRunnerServer) error {
	_e := util.NewExeErr("ConnectRunner", HOST)
	runner, exist := c.Nest.Runners[r.GetID()]
	if !exist {
		return status.Errorf(codes.InvalidArgument, _e.String("invalid request", "not active runner"))
	}
	outlet := make(chan pb.RunnerOutput)
	oi := len(runner.Outlets)
	runner.oLock.Lock()
	runner.Outlets = append(runner.Outlets, &outlet)
	runner.oLock.Unlock()
	defer func() {
		close(outlet)
		runner.oLock.Lock()
		runner.Outlets = append(runner.Outlets[:oi], runner.Outlets[oi+1:]...)
		runner.oLock.Unlock()
	}()
	runner.Log.Infof("session %v attached", oi)
	for {
		select {
		case msg := <-outlet:
			if err := stream.Send(&msg); err != nil {
				return status.Errorf(codes.Unavailable, err.Error())
			}
			c.Log.Tracef("message forwarded to session %v", oi)
		case <-runner.Ctx.Done():
			return nil
		}
	}
	return nil
}

func (c *Controller) JoinRunner(r *pb.RunnerUpdate, stream pb.Controller_JoinRunnerServer) error {
	_e := util.NewExeErr("JoinRunner", HOST)
	runner, exist := c.Nest.Runners[r.GetID()]
	if !exist {
		return status.Errorf(codes.InvalidArgument, _e.String("invalid request", "not active runner"))
	}
	// load previous data, sync runner states
	syncDone := false
	buf := []pb.RunnerOutput{}
	bLock := &sync.Mutex{}
	go func() {
		// load

		tt := []pb.RunnerOutput{pb.RunnerOutput{Msg: "MMMMMMMM"}}

		bLock.Lock()
		buf = append(tt, buf...)
		bLock.Unlock()
		// sync
		for _, msg := range buf {
			if err := stream.Send(&msg); err != nil {
				syncDone = true
				return
			}
		}
		syncDone = true
		return
	}()
	outlet := make(chan pb.RunnerOutput)
	oi := len(runner.Outlets)
	runner.oLock.Lock()
	runner.Outlets = append(runner.Outlets, &outlet)
	runner.oLock.Unlock()
	defer func() {
		close(outlet)
		runner.oLock.Lock()
		runner.Outlets = append(runner.Outlets[:oi], runner.Outlets[oi+1:]...)
		runner.oLock.Unlock()
	}()
	runner.Log.Infof("session %v attached", oi)
	for {
		select {
		case msg := <-outlet:
			if syncDone {
				if err := stream.Send(&msg); err != nil {
					return status.Errorf(codes.Unavailable, err.Error())
				}
				c.Log.Tracef("message forwarded to session %v", oi)
			} else {
				bLock.Lock()
				buf = append(buf, msg)
				bLock.Unlock()
			}
		case <-runner.Ctx.Done():
			return nil
		}
	}
	return nil
}

func (c *Controller) UpdateRunner(ctx context.Context, r *pb.RunnerUpdate) (*pb.OprStat, error) {
	_e := util.NewExeErr("UpdateRunner", HOST)
	runner, exist := c.Nest.Runners[r.GetID()]
	if !exist {
		return nil, status.Errorf(codes.InvalidArgument, _e.String("invalid request", "not active runner"))
	}
	result, err := runner.Wkr.UpdateRunner(runner.Ctx, r)
	if err != nil {
		runner.Log.WithError(err).Error("failed worker request, unable to update service")
		return nil, status.Errorf(codes.Unavailable, _e.String("unable to update service on worker", err))
	}
	return result, nil
}

// periodical worker Regist
func (c *Controller) RegistWorker(ctx context.Context, svr *pb.SvrStat) (*pb.OprStat, error) {
	_e := util.NewExeErr("RegistWorker", HOST)
	// registried worker
	if exist, _ := c.Pool.Find(svr.GetHost()); exist {
		return &pb.OprStat{State: "exist"}, nil
	}
	// regist
	// verify worker capabilities

	// connect worker gRPC server, use pre configured self-signing CA
	caPubKey, err := credentials.NewClientTLSFromFile("/appsrc/cert/wkr.cer", "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, _e.String("worker certification not exist", err))
	}
	// Set up TLS connection to the server
	wkrConn, err := grpc.Dial(svr.GetHost()+":50051", grpc.WithTransportCredentials(caPubKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, _e.String("unable to connect worker", err))
	}
	c.Pool.Add(&Worker{svr, wkrConn})
	c.Log.Infof("worker %s regist success, %v worker in pool %v", svr.GetHost(), len(c.Pool.Workers), c.Pool.Workers)
	return &pb.OprStat{State: "new"}, nil
}

func main() {
	mainCtx, mainCancel := context.WithCancel(context.Background())
	defer mainCancel()

	controller := Controller{
		Nest: Nest{make(map[string]*Runner), &sync.Mutex{}},
		Pool: Pool{[]*Worker{}, &sync.Mutex{}},
		Log:  log.WithField("owner", "rgw"),
	}
	// start garbage collector
	go controller.Nest.gc(mainCtx)
	// setup and run gRPC server
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		controller.Log.WithError(err).Fatal("gRPC server fail: unable to init tcp socket 50051")
	}
	// TLS
	grpcTLS, err := credentials.NewServerTLSFromFile("/appsrc/cert/rgw.cer", "/appsrc/cert/rgw.key")
	if err != nil {
		controller.Log.WithError(err).Fatal("gRPC server fail: invalid TLS keys")
	}
	// gRPC Controller server
	grpcSvr := grpc.NewServer(grpc.Creds(grpcTLS))
	pb.RegisterControllerServer(grpcSvr, &controller)
	go func() {
		controller.Log.Info("gRPC server start")
		controller.Log.Fatal(grpcSvr.Serve(grpcListener))
	}()

	hold := make(chan struct{})
	<-hold
}
