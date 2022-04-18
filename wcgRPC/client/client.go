package main

import (
	"context"
	"log"
	"wcgRPC/proto"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := proto.NewChatServiceClient(conn)

	response, err := c.WordCount(context.Background(), &proto.Input{
		Body: "Hii user",
	})

	if err != nil {
		log.Fatalf("Error when calling Input: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
