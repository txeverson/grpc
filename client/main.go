package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	user "gRPC/proto/gen"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8200", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	usr := user.NewUserServiceClient(conn)

	response, err := usr.GetUser(context.Background(), &user.UserRequest{Username: "txeverson"})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}

	log.Printf("Response from server: %v", response)
}
