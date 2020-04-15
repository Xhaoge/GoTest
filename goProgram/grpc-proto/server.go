package main

import (
	"context"
	"fmt"
	pb "./goProgram/grpc-proto/protos"
	"log"
	"net"

	"github.com/Xhaoge/sh/myhttp"
	"google.golang.org/grpc"
)

const (
	YUETUURL = "http://test-api.gloryholiday.com/yuetu/search"
	PORT     = "9091"
)

type SearchService struct{}

func buildYuetuReq(r *pb.YuetuSearchRequest) {
	myReq := myhttp.NewHttpSend(YUETUURL)
	fmt.Println("YuetuUrl: ", YUETUURL)
	fmt.Println("myReq: ", myReq)
	fmt.Println("YuetuSearchRequest: ", r)
}

func (s *SearchService) Search(ctx context.Context, r *pb.YuetuSearchRequest) (*pb.YuetuSearchResponse, error) {
	buildYuetuReq(r)
	return &pb.YuetuSearchResponse{}, nil
}

func main() {
	fmt.Println("grpc hello world.......")
	server := grpc.NewServer()
	pb.RegisterSHServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net .listen err： %v", err)
	}
	server.Serve(lis)
}
