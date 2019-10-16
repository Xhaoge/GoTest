package main

import ("math"
		"net"
		"net/rpc"
		"net/http")

type MathUtil struct{

}

// 该方法向外暴露，提供计算圆形面积的服务；
func (mu *MathUtil) CalculateCircleArea (req float32,resp *float32) error{
	*resp = math.pi * req * req          // 圆形的面积
	return nil                           //返回类型
}

func main(){
	//1 计划初始
}