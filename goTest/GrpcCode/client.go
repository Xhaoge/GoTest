package main

import (
	"GOlang/goTest/GrpcCode/message"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.DialHTTP("localhost:8089", grpc.WithInsecure)
	if err != nil {
		panic(err.Error())
	}
	defer client.Close()

	OrderServiceClient := message.NewOrderServiceServer(client)

	orderRequest := &message.OrderRequest{OrderId: "201907300003", TimeStamp: time.Now().Unix()}
	orderInfo, err := OrderServiceClient.GetOrderInfo(context.Background(), orderRequest)

	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
