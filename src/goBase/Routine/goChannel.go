package main

import (
	"fmt"
)

func main(){
	fmt.Println(" channel test 管道")
	var intChan chan int 
	intChan = make(chan int,3)
	//看看intChan是什么
	fmt.Printf("intChan = %v \n",&intChan)
	// 向管道写入数据;
	intChan<- 11
	num :=211
	intChan<- num
	intChan<- 20
	fmt.Printf("len = %v ;cap=%v \n",len(intChan),cap(intChan))
	// 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 = ",num2)
	fmt.Printf("len = %v ;cap=%v \n",len(intChan),cap(intChan))

}