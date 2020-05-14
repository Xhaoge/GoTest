package main

import (
	"fmt"
	"time"
)

// func main() {
//     go func() {
//     	proc()
//         // 1 在这里需要你写算法
//         // 2 要求每秒钟调用一次proc函数
//         // 3 要求程序不能退出
//     }()

//     select {}
// }

// func proc() {
//     panic("ok")
// }

func proc() {
	panic("ok")
}

func goproc(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%d panic err：%s\n", i, err)
		}
	}()
	proc()
}

func main() {
	i := 0
	go func() {
		for {
			ticker := time.NewTicker(time.Second * 1)
			defer ticker.Stop()
			select {
			case <-ticker.C:
				i++
				go goproc(i)
			}

		}
	}()
	select {}
}
