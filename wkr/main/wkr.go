package main

import (
	"context"
	"fmt"
	"net"
	"os"
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
	//HOST, _ = os.Hostname()
	HOST = os.Getenv("WKR")
	RGW  = os.Getenv("RGW") + ":50051"
)

func init() {
	// config package level default logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

type Runner struct {
	ID        string
	Task      string
	Ctx       context.Context
	Cancel    context.CancelFunc
	Out       chan string
	Log       *log.Entry
	Factor    int64
	paramLock *sync.Mutex
}

func (runner *Runner) Init() error {
	runner.Ctx, runner.Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	runner.Out = make(chan string)
	go func() {
		counter := int64(0)
		hold := time.NewTimer(time.Second)
		for {
			select {
			case <-hold.C:
				counter++
				select {
				case runner.Out <- fmt.Sprintf("---%s---%s---%v---\n", runner.ID, HOST, counter*runner.Factor):
				default:
				}
				hold = time.NewTimer(time.Second)
			case <-runner.Ctx.Done():
				return
			}
		}
	}()
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

type WorkerNode struct {
	pb.UnimplementedWorkerServer
	Nest
	Log *log.Entry
}

func (c *WorkerNode) Healtz(ctx context.Context, r *pb.Request) (*pb.SvrStat, error) {
	c.Log.Trace("health check response")
	return &pb.SvrStat{
		Host:      HOST,
		Load:      int64(len(c.Nest.Runners) * 100 / 3),
		Release:   "0.4.1",
		Instances: []*pb.SvcStat{},
	}, nil
}

func (c *WorkerNode) CreateRunner(ctx context.Context, r *pb.RunnerRequest) (*pb.SvcStat, error) {
	_e := util.NewExeErr("Create", HOST)
	if r.GetID() == "" {
		return nil, status.Errorf(codes.InvalidArgument, _e.String("invalid request", "missing runner id"))
	}
	if _, exist := c.Nest.Runners[r.GetID()]; exist {
		return nil, status.Errorf(codes.InvalidArgument, _e.String("invalid request", "duplicated runner"))
	}
	newRunner := Runner{
		ID:        r.GetID(),
		Task:      r.GetTask(),
		Log:       c.Log.WithField("runner", r.GetID()),
		Factor:    int64(1),
		paramLock: &sync.Mutex{},
	}
	if err := newRunner.Init(); err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	c.Nest.Add(&newRunner)
	c.Log.Infof("runner %s created", newRunner.ID)
	return &pb.SvcStat{ID: newRunner.ID, Healtz: true, State: "create"}, nil
}

func (c *WorkerNode) ConnectRunner(r *pb.RunnerUpdate, stream pb.Worker_ConnectRunnerServer) error {
	runner := c.Nest.Runners[r.GetID()]
	defer runner.Cancel()
	c.Log.Infof("runner %s connected", runner.ID)
	for {
		select {
		case msg := <-runner.Out:
			if err := stream.Send(&pb.RunnerOutput{Msg: msg}); err != nil {
				return status.Errorf(codes.Unavailable, err.Error())
			}
			c.Log.Tracef("message sent: %v", msg)
		case <-runner.Ctx.Done():
			return nil
		}
	}
	return nil
}

func (c *WorkerNode) UpdateRunner(ctx context.Context, r *pb.RunnerUpdate) (*pb.OprStat, error) {
	_e := util.NewExeErr("UpdateRunner", HOST)
	runner, exist := c.Nest.Runners[r.GetID()]
	if !exist {
		return nil, status.Errorf(codes.InvalidArgument, _e.String("invalid request", "not active runner"))
	}
	runner.paramLock.Lock()
	runner.Factor = r.GetFactor()
	runner.paramLock.Unlock()
	return &pb.OprStat{State: "done"}, nil
}

func main() {
	sLog := log.WithField("owner", HOST)
	// regist to rgw
	// config gRPC server, using rgw self-signing CA
	caPubKey, err := credentials.NewClientTLSFromFile("/appsrc/cert/rgw.cer", "")
	if err != nil {
		sLog.WithError(err).Fatal("unable to import rgw certificate")
	}
	// Set up TLS connection to the server
	rgwConn, err := grpc.Dial(RGW, grpc.WithTransportCredentials(caPubKey))
	if err != nil {
		sLog.WithError(err).Fatal("unable to connect rgw server")
	}
	defer rgwConn.Close()
	rgw := pb.NewControllerClient(rgwConn)
	gCtx, gCancel := context.WithCancel(context.Background())
	defer gCancel()

	go func() {
		for {
			select {
			case <-gCtx.Done():
				return
			default:
				if state, err := rgw.RegistWorker(gCtx, &pb.SvrStat{
					Host:      HOST,
					Load:      0,
					Release:   "0.4.1",
					Instances: []*pb.SvcStat{},
				}); err != nil {
					sLog.WithError(err).Warn("worker registration fail")
				} else {
					switch state.GetState() {
					case "new":
						sLog.Info("registration success")
					case "exist":
						//
					default:
						//
					}
				}
			}
			time.Sleep(time.Second * 5)
		}
	}()

	// setup and run gRPC server
	mainCtx, mainCancel := context.WithCancel(context.Background())
	defer mainCancel()

	node := WorkerNode{
		Nest: Nest{make(map[string]*Runner), &sync.Mutex{}},
		Log:  sLog,
	}
	// start garbage collector
	go node.Nest.gc(mainCtx)

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		sLog.WithError(err).Fatal("gRPC server fail: unable to init tcp socket 50051")
	}
	// TLS
	grpcTLS, err := credentials.NewServerTLSFromFile("/appsrc/cert/wkr.cer", "/appsrc/cert/wkr.key")
	if err != nil {
		sLog.WithError(err).Fatal("gRPC server fail: invalid TLS keys")
	}

	grpcSvr := grpc.NewServer(grpc.Creds(grpcTLS))

	// gRPC Controller server
	pb.RegisterWorkerServer(grpcSvr, &node)

	go func() {
		sLog.Info("gRPC server start")
		sLog.Fatal(grpcSvr.Serve(grpcListener))
	}()

	hold := make(chan struct{})
	<-hold
}
