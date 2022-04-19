package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// s := proto.Server{}

	// grpcServer := grpc.NewServer()

	//proto.RegisterChatServiceServer(grpcServer, &s)

	// proto.RegisterChatServiceServer(grpcServer,&s)

	fmt.Println(lis)
}
