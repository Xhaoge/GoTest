package main

import (
	"GOlang/goTest/GrpcCode/message"
	"errors"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

//订单服务
type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"201907300001": {OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907300002": {OrderId: "201907300002", OrderName: "零食", OrderStatus: "已付款"},
		"201907300003": {OrderId: "201907300003", OrderName: "食品", OrderStatus: "未付款"},
	}

	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
			return &result, nil
		} else {
			return errors.New("server error")
		}
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(lis, nil)
}
