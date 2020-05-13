package main

import (
	"fmt"
	"sync"
)

/*
go提供了sync包和channel来解决协程同步和通讯。新手对channel通道操作起来更容易产生死锁，如果时缓冲的channel还要考虑channel放入和取出数据的速率问题。
从字面就可以理解，sync.WaitGroup是等待一组协程结束。它实现了一个类似任务队列的结构，你可以向队列中加入任务，任务完成后就把任务从队列中移除，如果队列中的任务没有全部完成，队列就会触发阻塞以阻止程序继续运行。
sync.WaitGroup只有3个方法，Add()，Done()，Wait()。 其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
简单示例如下：*/

var wg sync.WaitGroup

func test(n int) {
	fmt.Printf("i = %d\n", n)
	wg.Done()
}

func main() {
	fmt.Println("xxxxxx")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("exit")
}
