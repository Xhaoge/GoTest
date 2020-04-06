package main

import (
	"fmt"
	"net" // 做网络socket开发时，net包含所有方法和函数
)

func main() {
	fmt.Println("服务器开始监听。。。。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭该接口；
	for {
		fmt.Println("等待客户端来连接。。。。")
		conn, err := listen.Accept() //等待客户端来连接；
		if err != nil {
			fmt.Println("accept() err=", err)
		} else {
			fmt.Println("accept() suc son=%v\n", conn)
		}
	}

	fmt.Println("listen suc=%v", listen)

}
