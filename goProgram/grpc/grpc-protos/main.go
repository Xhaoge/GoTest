package main

import (
	"fmt"
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "Pro_golang/Golang/grpc-protos/search"
)

const (
	PORT = ":50001"
)

type server struct{}

func (s *server)SayHello(ctx context.Context,in *pb.HelloRequest) (*pb.HelloReply,error) {
	log.Println("request: ",in.Name,in.Age,in.Work)
	return &pb.HelloReply{Message:"good luck to you",Work:"nima",Age:17},nil
}

func main(){
	fmt.Println("this is test for grpc")
	lis,err := net.Listen("tcp",PORT)

	if err != nil {
		log.Fatalf("fail to lister port :%v",err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServer(s,&server{})
	log.Println("grpc 服务已经开启；请开始你的表演；")
	s.Serve(lis)
}
