package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/ra9dev/go-template/pb"
)

func main() {
	conns, err := grpc.Dial(":82", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewGreeterClient(conns)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("resp: %v", resp)
}
