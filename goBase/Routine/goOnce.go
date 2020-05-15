package main

import (
    "fmt"
    "sync"
)

/*
https://studygolang.com/articles/28570?fr=sidebar*/

func onceFunc(i int) {
    fmt.Printf("goroutine %d run\n", i)
}

func main() {
    ch := make(chan struct{}, 10)
    var once sync.Once

    for i:=0;i<10;i++ {
		fmt.Println("i 的打印：",i)
        go func() {
            once.Do(func() {
                onceFunc(i)
            })
            ch <- struct{}{}
        }()
    }

    for i:=0;i<10;i++ {
        <- ch
    }
}