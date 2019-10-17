package main

import (
	"google.golang.org/grpc"
	"GOlang/goTest/socket/GrpcCode/message"
	"context"
	"fmt"
	"time"
)

func main() {
	client, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer client.Close()

	OrderServiceClient := message.NewOrderServiceClient(client)

	orderRequest := &message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()}
	orderInf, err := OrderServiceClient.GetOrderInfo(context.Background(), orderRequest)

	fmt.Println("nima,where")
	fmt.Println(orderInf.GetOrderId())
	if orderInf != nil {
		fmt.Println(orderInf.GetOrderId())
		fmt.Println(orderInf.GetOrderName())
		fmt.Println(orderInf.GetOrderStatus())
	}
}
