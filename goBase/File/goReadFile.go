package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
)

func main() {
	/*
		fileinfo: 文件信息
			interface
				Name():文件名
				Size()：文件大小，字节为单位；
				IsDir():是否是目录；
				ModTime():修改时间；
				Mode()：权限

	*/
	//获取当前路径；
	fmt.Println("read file test!")
	dir, _ := os.Getwd()
	fmt.Println("当前路径是： ", dir)
	file, err := os.Open("aa.txt")
	if err != nil {
		fmt.Println("open file err= ", err)
		return
	}
	defer file.Close()
	// 输出文件，看看文件是什么；
	fmt.Printf("file = %v\n",file)
	// 带缓冲的reader
	reader := bufio.NewReader(file)
	// 循环的读取文件的内容
	for {
		str,err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF{ // io.EOF表示文件的末尾
			break
		} 
		fmt.Print(str)
	}
	fmt.Println("缓冲读取文件结束。。。。。。")
	f := "D:/program/go/golang/src/goBase/File/bb.md"
	content,err := ioutil.ReadFile(f)
	if err != nil{
		fmt.Printf("read file err=%v",err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v\n",content)  //[]byte
	fmt.Printf("%v",string(content))  //转成string

	
}
