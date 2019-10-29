package main

import (
	"fmt"
	"os"
	"path"
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
	dir, _ := os.Getwd()
	s := path.Join(dir, "..")
	fmt.Println("当前路径是： ", s)
	fileInfo, err := os.Stat("aa.txt")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	// filename
	fmt.Println(fileInfo.Name())
	// file size
	fmt.Println(fileInfo.Size())
}
