package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	)

// 自己编写一个函数，接受2个文件路径；srcfile dstfile

func CopyFile(dstfile string,srcfile string)(written int64,err error){
	src,err := os.Open(srcfile)
	if err !=nil{
		fmt.Println("open file err=",err)
	}
	defer src.Close()
	// 通过src，获取到reader
	reader := bufio.NewReader(src)
	//打开dst file
	dst,err := os.OpenFile(dstfile,os.O_WRONLY | os.O_CREATE,0666)
	if err !=nil{
		fmt.Println("open err=",err)
	}
	defer dst.Close()
	//通过det 获取到writer
	writer := bufio.NewWriter(dst)
	return io.Copy(writer,reader)
}

func main(){
	fmt.Println("copy file gif test!")
	srcfile :="D:/program/file/mm.jpg"
	dstfile :="D:/program/go/golang/src/goBase/File/mm.jpg"
	CopyFile(dstfile,srcfile)
	fmt.Println("copy finish!")
}