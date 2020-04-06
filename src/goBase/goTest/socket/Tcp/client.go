package main

import ("fmt"
		"net/rpc")

//客户端逻辑实现
func main(){
	client,err := rpc.DialHTTP("tcp","localhost:8082")
	if err != nil{
		panic(err.Error())
	}
	var req float32 
	req = 1
	var resp *float32
	//同步的调用方式；
	err = client.Call("MathUtil.CalculateCircleArea",req,&resp)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(*resp)

}