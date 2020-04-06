package main

import (
	"GOlang/goTest/socket/RpcCode/message"
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

//订单服务
type OrderService struct {
}

func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
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
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	OrderService := new(OrderService)
	rpc.Register(OrderService)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
