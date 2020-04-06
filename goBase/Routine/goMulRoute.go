package main 

import (
	"fmt"
	"time"	
)

func putNum(intChan chan int) {
	for i:=1; i<=8000; i++ {
		intChan<- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool){
	var flag bool
	for {
		time.Sleep(time.Millisecond*10)
		num, ok := <- intChan
		if !ok {
			break
		}
		flag = true
		// 判断num是否为素数
		for i:=2; i < num ; i++ {
			if num % i == 0{
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num
		}
	}
	fmt.Println("协程取不到数据，退出")
	exitChan<- true
}




func main() {
	fmt.Println("多个线程跑，辨别出素数")
	intChan := make(chan int, 1000)
	primeChan := make(chan int,2000)
	// 标是退出的管道
	exitChan := make(chan bool, 4)

	//开启协程，向intchan 里放入1-8000个数
	go putNum(intChan)
	// 开启4个携程，从intchan 取出数据，判断是否为素数；如果是就放primeChan
	for i:=0 ;i<4; i++ {
		go primeNum(intChan,primeChan,exitChan)
	}
	//等待四个协程全部完成，再主程序退出
	go func(){
		for i:=0 ;i<4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 遍历我们的primeNum，结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("打印素数为：",res)
	}
	
}