package main

import ("net/rpc"
		"net"
		"net/http"
		"math")

//数学计算
type MathUtil struct{

}

//该方法向外暴露，提供计算圆形面积的服务；
func (mu *MathUtil) CalculateCircleArea(req float32,resp *float32) error {
	*resp = math.Pi * req *req
	return nil
}

//main 方法
func main(){
	// 初始化指针数据类型
	mathUtil := new(MathUtil)
	// 调用net/rpc包的功能将服务对象进行注册；
	err := rpc.Register(mathUtil)
	if err != nil{
		panic(err.Error())
	}
	//通过该函数把mathUtil中提供的注册到http协议上，方便调用者可以使用http的方式进行数据传递；
	rpc.HandleHTTP()
	//在特定的接口进行监听；
	listen,err := net.Listen("tcp",":8082")
	if err != nil{
		panic(err.Error())
	}
	http.Serve(listen,nil)
}