package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/vinigracindo/gogrpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	AddUserStream(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserStream(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	responseStream, err := client.AddUserStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not read stream: %v", err)
		}
		fmt.Println("Status:", stream.Status, "User:", stream.User)
	}
}
