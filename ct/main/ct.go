package main

import (
	"context"
	"io"
	"os"
	"strconv"
	"time"

	pb "protobuf/synergy"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	log "github.com/sirupsen/logrus"
)

var (
	// local host name
	HOST, _   = os.Hostname()
	RGW       = os.Getenv("RGW") + ":50051"
	FACTOR, _ = strconv.ParseInt(os.Getenv("FACTOR"), 10, 64)
)

func init() {
	// config package level default logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {
	sLog := log.WithField("owner", HOST)
	// config gRPC server
	// pre configured self-signing CA
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

	runner, err := rgw.CreateRunner(gCtx, &pb.RunnerRequest{Task: "ct1"})
	if err != nil {
		sLog.WithError(err).Error("failed runner request")
		return
	}
	sLog.Infof("got runner %v", runner)

	stream, err := rgw.ConnectRunner(gCtx, &pb.RunnerUpdate{ID: runner.GetID()})
	if err != nil {
		sLog.WithError(err).Error("failed runner request")
		return
	}
	go func() {
		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				sLog.Infof("streaming end: %v", err)
				return
			}
			if err != nil {
				sLog.Warnf("streaming error: %v", err)
				return
			}
			log.Println(feature)
		}
	}()

	time.Sleep(time.Second * 5)
	// stop
	gCancel()

	time.Sleep(time.Second * 5)
	// join
	gCtx, gCancel = context.WithCancel(context.Background())
	stream, err = rgw.JoinRunner(gCtx, &pb.RunnerUpdate{ID: runner.GetID()})
	if err != nil {
		sLog.WithError(err).Error("failed runner request")
		return
	}
	go func() {
		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				sLog.Infof("streaming end: %v", err)
				return
			}
			if err != nil {
				sLog.Warnf("streaming error: %v", err)
				return
			}
			log.Println(feature)
		}
	}()

	time.Sleep(time.Second * 5)
	// update

	_, err = rgw.UpdateRunner(gCtx, &pb.RunnerUpdate{ID: runner.GetID(), Factor: FACTOR})
	if err != nil {
		sLog.WithError(err).Error("failed runner request")
		return
	}

	hold := make(chan struct{})
	<-hold
}
