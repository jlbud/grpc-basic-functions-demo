package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	pb "grpc-test/proto"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

const address = "127.0.0.1:8001"

func TestCli(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)

	name := "xiao ming"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Hello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("client response: %s, %s", r.Message, r.Bot.Run)
}

func TestClientStreamCli(t *testing.T) {
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	client := pb.NewHelloWorldClient(conn)

	cli, _ := client.ClientStreamHello(context.Background())

	for n := 1; n < 10; n++ {
		s := pb.HelloStreamRequest{
			Name: strconv.Itoa(n),
		}
		err := cli.Send(&s)
		if err != nil {
			t.Error(err)
			return
		}
	}
	resp, err := cli.CloseAndRecv()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestDoubleStreamCli(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
