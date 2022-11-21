package main

import (
	"context"
	. "github.com/LittleMikle/grpc_go/gen/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type testApiServer struct {
	UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *ResponseRequest) (*ResponseRequest, error) {
	return req, nil
}

func main() {
	go func() {
		// mux
		mux := runtime.NewServeMux()

		//register handler
		RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})

		//launch http server
		log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	}()

	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
