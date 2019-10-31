package main

import (
	"fmt"
)

type Cat struct{
	name string
	age int
}

func writedata(a chan int){
	for i:=1;i<=50;i++{
		a <-i
		fmt.Println("write =",i)
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

	// var allchan chan interface{}

	allchan := make(chan interface{},3)
	allchan <- 10
	allchan <- "tome jack"
	cat := Cat{"小花猫哦",4}
	allchan <- cat
	// 我们希望或得到管道中的第三个元素，先放出前2个
	<-allchan
	<-allchan
	newcat := <-allchan
	fmt.Printf("newcat = %T,NERCAT=%v\n ",newcat,newcat)
	// 使用类型断言
	a := newcat.(Cat)
	fmt.Printf("newcat.name = %v",a.name)

	intchan1 := make(chan int,50)
	exitchan := make(chan bool,1)
	
	go writedata(intchan1)
	go readdata(intchan1,exitchan)
}