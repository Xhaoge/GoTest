package main

import (
	"fmt"
	"time"
)

type Cat struct {
	Name 	string
	Age 	int
}

func writedata(a chan int){
	fmt.Println("write")
	for i:=1;i<=50;i++{
		a <-i
		fmt.Println("write =",i)
		time.Sleep(time.Second)
	}
	close(a)
}

func readdata(b chan int,c chan bool){
	for {
		v,ok := <-b
		if !ok{
			break
		}
		fmt.Printf("read = %v\n",v)
	}
	c <- true
	close(c)
}

func main(){
	fmt.Println(" channel test 管道")
	//var allChan	chan interface{}
	allChan := make(chan interface{}, 3)
	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"x浩哥", 4}
	allChan<- cat
	// 希望获取管道中的第三个元素，则需要将前2个推出；
	<- allChan
	<- allChan

	newCat := <-allChan  // 从管道中取出的类型是interface
	// 需要使用类型断言取出
	a := newCat.(Cat)
	fmt.Println("newCat name：",a.Name)
	// intchan1 := make(chan int,50)
	// exitchan := make(chan bool,1)
	
	// go writedata(intchan1)
	// go readdata(intchan1,exitchan)
	// for {
	// 	_,ok := <-exitchan
	// 	if !ok{
	// 		break
	// 	}
	// }
}