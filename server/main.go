package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Elrathor/Larian-Adventure-Coop-Extender/lace"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement lace.ExchangeServer.
type server struct {
	pb.UnimplementedExchangeServer
}

func (s *server) SendCommand(ctx context.Context, in *pb.SendCommandRequest) (*pb.SendCommandReply, error) {
	return &pb.SendCommandReply{Success: true}, nil
}

func (s *server) GetCommand(ctx context.Context, in *pb.GetCommandRequest) (*pb.GetCommandReply, error) {
	return &pb.GetCommandReply{Command: "SAVE"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExchangeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
