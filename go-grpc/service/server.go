package main

import (
	"context"
	"fmt"
	pb "godie/go-grpc/protos"
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

func buildYuetuReq(r *pb.YuetuSearchRequest)([]byte, error) {
	myReq := myhttp.NewHttpSend(YUETUURL)
	fmt.Println("YuetuUrl: ", YUETUURL,":",PORT)
	fmt.Println("myReq: ", myReq)
	fmt.Println("YuetuSearchRequest: ", r)
	req := map[string]string{
		"Cid":r.BaseRequest.Cid,
		"FromCity":r.Trip.DepartureCode,
		"ToCity":r.Trip.ArrivalCode,
		"Cabin":string(r.Cabin),
		"FromDate":"20200627",
		"TripType":"1",
	}
	myReq.SetSendType("json")
	myReq.SetBody(req)
	fmt.Println("%v: ",myReq)
	res,err := myReq.Post()
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return res,nil
}

func (s *SearchService) Search(ctx context.Context, r *pb.YuetuSearchRequest) (*pb.YuetuSearchResponse, error) {
	res ,err := buildYuetuReq(r)
	fmt.Println("res err: ",err)
	fmt.Println("res:  ",string(res))
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
