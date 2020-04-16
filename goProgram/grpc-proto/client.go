package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "godie/goProgram/grpc-proto/protos"
)

const PORT = "9091"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSHServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.YuetuSearchRequest{
		BaseRequest: &pb.SimpleRequest{
			Cid: "qunarytb",
			TraceId: "2354435-jgidg",
		},
		Trip:&pb.Trip{
			DepartureCode:"CNX",
			ArrivalCode:"BKK",
			DepartureDate:"2020-07-20T16:00:00Z",
		},
		AdultNum:1,
		ChildNum: 0,
		InfantNum:0,
		BypassCache:false,
		GodPerspective:false,
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp)
}