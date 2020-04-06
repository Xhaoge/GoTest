package main

import (
	"fmt"
	"bufio"
	"os"
	)

func main(){
	fmt.Println("write file test!")
	file,err := os.OpenFile("bb.md",os.O_WRONLY | os.O_APPEND,0666)
	if err != nil{
		fmt.Println("open file err=",err)
	}

	// 准备写入5句 “hello world”
	str := "hello world!\n"
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	// 因为write 是带缓存的，因此在调用write 方法时，其实内容是先写入缓存的；
	writer.Flush()
}