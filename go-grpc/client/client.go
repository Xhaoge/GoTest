package main

import (
	"context"
	"fmt"
	pb "godie/go-grpc/protos"
	"log"

	"google.golang.org/grpc"
)

const PORT = "9091"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSHServiceClient(conn)
	req := &pb.YuetuSearchRequest{
		BaseRequest: &pb.SimpleRequest{
			Cid: "qunarytb",
			TraceId: "2354435-jgidg",
		},
		Trip:[]*pb.Trip{{
			DepartureCode: "CNX",
			ArrivalCode:   "BKK",
		},
		},
		Cabin: pb.CabinClass_E,
		AdultNum:1,
		ChildNum: 0,
		InfantNum:0,
		BypassCache:false,
		GodPerspective:false,
		TargetProviders: []string{"mondee"},
	}
	fmt.Println("req :",req.TargetProviders)
	resp, err := client.Search(context.Background(), req)

	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Println("resp: ",resp)
	//log.Printf("resp: %s", resp.GetRouting()[0].GetData())
	//log.Printf("resp: %s ",resp.GetBaseResponse())
	//log.Printf("resp: %f ",resp.GetRouting()[0].GetAdultPrice())
}