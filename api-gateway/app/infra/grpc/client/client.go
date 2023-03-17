package client

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"api-gateway/app/infra/grpc/proto-gen/message"
)

func SayHello(name string) (string, error) {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	c := message.NewExampleServiceClient(conn)
	req := &message.HelloRequest{
		Name: name,
	}

	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		return "", err
	}

	log.Printf("Received response from server: %s", resp.Message)
	return resp.Message, nil
}
