package main

import (
	user "gRPC/proto/gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	usr := Server{}
	user.RegisterUserServiceServer(grpcServer, &usr)

	log.Println("Listening on Port: 8200!")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
