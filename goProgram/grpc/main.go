package main

import (
	"fmt"
	_"Pro_golang/Golang/config"
	//"Pro_golang/Golang/TestCase/SearchCase"
	//"Pro_golang/Golang/utils"
	"net/http"
	"log"
)

func sayHelloGolang(w http.ResponseWriter,r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的；
	fmt.Println("path:",r.URL.Path)
	w.Write([]byte("Hello Golang"))
}

func main(){
	fmt.Println("hello world.......")
	http.HandleFunc("/",sayHelloGolang)  // 设置访问路由
	err := http.ListenAndServe(":8080",nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}

