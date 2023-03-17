package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"api-gateway/app/infra/grpc/proto-gen/message"
)

type exampleServer struct {
}

func (s *exampleServer) SayHello(ctx context.Context, req *message.HelloRequest) (*message.HelloResponse, error) {
	log.Printf("Received request from client: %s", req.Name)
	resp := &message.HelloResponse{
		Message: "Hello " + req.Name,
	}
	return resp, nil
}

func RunServer() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	message.RegisterExampleServiceServer(s, &exampleServer{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
