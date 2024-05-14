package main

import (
	"context"
	"log"
	"net"

	"gRPC/weblog"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	weblog.UnimplementedWebLoggerServer
}

// The actual log function
func (s *server) Log(ctx context.Context, wr *weblog.WebRequest) (*empty.Empty, error) {

	log.Println(wr.String())

	return &emptypb.Empty{}, nil

}

func main() {

	//Step 1 : Create a new Server
	grpcServer := grpc.NewServer()

	//Step 2: Register rpc services
	weblog.RegisterWebLoggerServer(grpcServer, &server{})

	//Step 3: start listening on hostAddr
	l, e := net.Listen("tcp", "localhost:8000")
	if e != nil {
		log.Fatalf("Failed to listen")
	}

	//Step 4: Start the RPC server
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}
