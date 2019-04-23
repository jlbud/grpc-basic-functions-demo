package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-test/proto"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)

	stream, err := c.DoubleStreamHello(context.Background())
	if err != nil {
		log.Printf("failed to call: %v", err)
		return
	}

	var i int64
	for {
		stream.Send(&pb.HelloStreamRequest{Name: "xiao hong"})
		if err != nil {
			log.Printf("failed to send: %v", err)
			break
		}
		reply, err := stream.Recv()
		if err != nil {
			log.Printf("failed to recv: %v", err)
			break
		}
		log.Printf("this client, server reply: %s", reply.Name)
		i++
		time.Sleep(1 * time.Second)
		if i == 10 {
			break
		}
	}
}
