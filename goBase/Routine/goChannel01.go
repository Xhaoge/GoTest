package main

import (
	"fmt"
	"time"
)

func main() {
	var c1 chan string = make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result"
	}()
	select {
	case f := <-c1:
		fmt.Println(f)
		// default:
		// 	fmt.Println("xxxx")
	}
	// fmt.Println("c1 is :", <-c1)
}
