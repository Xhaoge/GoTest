package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "goGrpc/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8089"
)

// server is used to implement helloworld.greeterserver
type server struct {
}

// sayhello implements helloworld.greeterserver
func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	printmyself()
	rr = pb.HelloReply{Message: "Hello " + r.Name+"  i konw your work:"+r.Work}
	return rr, nil
}

func printmyself() {
	fmt.Println("这是个什么鬼.......")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServer(s, &server{})
	// Register reflection service on grpc server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to seve :%v", err)
	}
}
