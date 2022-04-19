package proto

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) WordCount(ctx context.Context, in *Input) (*Input, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Input{Body: "Hello From the Server!"}, nil
}
