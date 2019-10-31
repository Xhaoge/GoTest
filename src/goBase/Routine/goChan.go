package main

import (
	"fmt"
	"time"
)

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
	intchan1 := make(chan int,50)
	exitchan := make(chan bool,1)
	
	go writedata(intchan1)
	go readdata(intchan1,exitchan)
	for {
		_,ok := <-exitchan
		if !ok{
			break
		}
	}
}