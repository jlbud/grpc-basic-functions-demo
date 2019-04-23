package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-test/proto"
)

const (
	port = ":5050"
)

type server struct{}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("has request: %s call you", in.Name)
	return &pb.HelloResponse{
		Message: fmt.Sprintf("hello %s ", in.Name),
		Bot: &pb.Bot{
			Run: "bot run",
		},
	}, nil
}

func (s *server) DoubleStreamHello(stream pb.HelloWorld_DoubleStreamHelloServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		log.Printf("server has request, client call: %s", resp.Name)
		stream.Send(&pb.HelloStreamResponse{Name: resp.Name})
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
