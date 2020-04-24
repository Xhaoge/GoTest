package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "godie/go-grpc/protos"
	"log"
	"net"
	"reflect"

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
		"FromCity":r.Trip[0].DepartureCode,
		"ToCity":r.Trip[0].ArrivalCode,
		"Cabin":"E",
		"FromDate":"20200628",
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
	fmt.Println("type res: ",reflect.TypeOf(&res))
	var resp pb.YuetuSearchResponse
	err := json.Unmarshal(res,&resp)
	if err != nil {
		fmt.Println("[]byte to struct err: ",err)
	}
	fmt.Println("YuetuSearchResponse string:  ",string(res))
	fmt.Println("YuetuSearchResponse baseresponse:  ",resp.BaseResponse)
	fmt.Println("YuetuSearchResponse routing:  ",resp.Routing)
	fmt.Println("YuetuSearchResponse:  ",resp)
	return &resp,nil
	//return &pb.YuetuSearchResponse{
	//	BaseResponse:&pb.SimpleResponse{
	//		Status:200,
	//		Message:"success",
	//		Cid:"yuetu",
	//		TraceId:"jwoejoag-jiejg",
	//		Pid:"mondee",
	//	},
	//	SessionId:"24526345662",
	//}, nil
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
