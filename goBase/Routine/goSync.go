package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func testFunc() {
	m.Lock()
	fmt.Println("i am going")
	time.Sleep(time.Second * 2)
	fmt.Println("hello world")
	m.Unlock()
}

func main() {
	m.Lock()
	go testFunc()
	time.Sleep(time.Second)
	m.Unlock()
	fmt.Println("finish...")
	for {
		time.Sleep(time.Second)
		fmt.Println("暂停....")
	}
}
