package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	example "github.com/ra9dev/go-template/pb"
)

type ExampleService struct {
	example.UnimplementedGreeterServer
}

func NewExampleService() ExampleService {
	return ExampleService{}
}

func (s ExampleService) SayHello(ctx context.Context, _ *emptypb.Empty) (*example.HelloReply, error) {
	return &example.HelloReply{Message: "Hello, world!"}, nil
}
