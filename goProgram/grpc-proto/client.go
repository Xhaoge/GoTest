package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/EDDYCJY/go-grpc-example/proto"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSHServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.YuetuSearchRequest{
		BaseRequest: &SimpleRequest{
			Cid: "tongchengyt",
			Trace_id: "2354435-jgidg",
		},
		Trip:&Trip{
			DepartureCode:"SIN",
			ArrivalCode:"SYD",
			DepartureDate:"2020-07-20T16:00:00Z",
		},
		Cabin:"E",
		AduleNum:1,
		ChildNum: 0,
		InfantNum:0,
		BypassCache:false,
		GodPerspective:false,
		TargetProviders:[],
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}