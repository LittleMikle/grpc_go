package main

import (
	"google.golang.org/grpc"
	pb "github.com/LittleMikle/grpc_go.git/gen/proto"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer()
}
