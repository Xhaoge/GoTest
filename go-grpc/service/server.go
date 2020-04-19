package main

import (
	"context"
	"encoding/json"
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
		"FromDate":r.Trip.DepartureDate,
		"TripType":"1",
	}
	myReq.SetSendType("json")
	myReq.SetBody(req)
	res,err := myReq.Post()
	if err != nil {
		return nil,err
	}
	return res,nil
}


func (s *SearchService) Search(ctx context.Context, r *pb.YuetuSearchRequest) (*pb.YuetuSearchResponse, error) {
	res ,_ := buildYuetuReq(r)
	var resp pb.YuetuSearchResponse
	err := json.Unmarshal(res,&resp)
	if err != nil {
		fmt.Println("sring to struct err: ",err)
	}
	fmt.Println("YuetuSearchResponse string:  ",string(res))
	return &resp, nil
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
