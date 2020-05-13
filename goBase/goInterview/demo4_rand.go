package main

import (
	"fmt"
	"math/rand"
)

/* 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel
中读取数字并打印到标准输出。最终输出五个随机数。
*/

func getRand(ch chan int) {
	n := rand.Intn(100)
	ch <- n
}

func printRand(ch chan int) {
	fmt.Printf("rand num：%d\n", <-ch)
}

// 方案一
// func main() {
// 	fmt.Println("开始.....")
// 	testChan := make(chan int)
// 	for i := 0; i < 5; i++ {
// 		go getRand(testChan)
// 		go printRand(testChan)
// 	}
// 	time.Sleep(time.Second * 2)
// 	close(testChan)
// 	// select {}
// }

//方案二

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case ch <- rand.Intn(100):
			case <-done:
				return
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Rand number ", <-ch)
		}
		done <- true
		return
	}()
	<-done
}
