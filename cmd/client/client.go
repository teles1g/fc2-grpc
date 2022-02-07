package main

import (
	"context"
	"fmt"
	"log"
	"github.com/teles1g/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connecto to gRPC: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Lucas Teles",
		Email: "dev.teles@outlook.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}