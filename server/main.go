package main

import (
	"container/list"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Elrathor/Larian-Adventure-Coop-Extender/lace"
	"google.golang.org/grpc"
)

var (
	port         = flag.Int("port", 50051, "The server port")
	commandQueue = list.New()
)

// server is used to implement lace.ExchangeServer.
type server struct {
	pb.UnimplementedExchangeServer
}

func (s *server) SendCommand(ctx context.Context, in *pb.SendCommandRequest) (*pb.SendCommandReply, error) {
	if in.GetCommand() == "SAVE" {
		commandQueue.PushBack(in.GetCommand())
		return &pb.SendCommandReply{Success: true}, nil
	} else {
		return &pb.SendCommandReply{Success: false}, nil
	}
}

func (s *server) GetCommand(ctx context.Context, in *pb.GetCommandRequest) (*pb.GetCommandReply, error) {
	if commandQueue.Len() > 0 {
		cmd := commandQueue.Front()
		commandQueue.Remove(cmd)
		return &pb.GetCommandReply{Command: cmd.Value.(string)}, nil
	} else {
		return &pb.GetCommandReply{Command: "NOOP"}, nil
	}

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
