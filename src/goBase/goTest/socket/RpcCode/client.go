package main

import (
	"GOlang/goTest/socket/RpcCode/message"
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8083")
	if err != nil {
		panic(err.Error())
	}

	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201907300001", TimeStamp: timeStamp}

	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*response)
}
