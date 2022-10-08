package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	example "github.com/ra9dev/go-template/pb"
	"github.com/ra9dev/go-template/pkg/tracing"
)

type ExampleService struct {
	example.UnimplementedGreeterServer
}

func NewExampleService() ExampleService {
	return ExampleService{}
}

func (s ExampleService) SayHello(ctx context.Context, _ *emptypb.Empty) (*example.HelloReply, error) {
	ctx, span := tracing.StartGRPCTrace(ctx, "exampleService.SayHello")
	defer span.End()

	exampleInternalBusinessLogicCall(ctx)

	return &example.HelloReply{Message: "Hello, world!"}, nil
}

// ExampleInternalBusinessLogicCall is an example of passing ctx and span to internal business logic.
func exampleInternalBusinessLogicCall(ctx context.Context) {
	_, span := tracing.SpanFromContext(ctx, "grpc", "someService.Hi")
	defer span.End()
}
