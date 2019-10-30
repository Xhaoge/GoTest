package main

import (
	"fmt"
	"time"
	"strconv"
	"runtime"
	"sync"
)

// 编写一个函数，每隔一秒输出“hello world
func test(){
	for i :=1;i<10;i++{
		fmt.Println("hello world "+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 编写一个函数，计算各个数的阶乘，并放入到map中；启动多个协程，放入map中；
var (
	mymap = make(map[int]int,10)
	// 声明一个全局的互斥锁;lock 是互斥锁，sync 是包，Mutex 是互斥； syncchornized 是同步
	lock sync.Mutex
)
func channelTest(n int){
	lock.Lock()
	res := 1
	for i:=1; i<=n; i++{
		res *=i
	}
	// 这里将res 放入到map中
	mymap[n] =res
	lock.Unlock()
}

func main(){
	fmt.Println("go routine test!")
	go test() // 开启一个协程；

	for i :=1;i<3;i++{
		fmt.Println("main hello world "+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	// runtime 获取电脑的cpu数目；
	cpuNum := runtime.NumCPU()
	fmt.Println("电脑cpu数目：",cpuNum)
	// runtime.GOMAXPROCS() 设置cpu使用数目；
	for i :=1;i<=20;i++{
		go channelTest(i)
	}
	lock.Lock()
	// 这里我们输出结果，变量这个结果；
	for i,v := range mymap{
		fmt.Printf("%d 的阶乘是：%v\n",i,v)
	}
	lock.Unlock()

}