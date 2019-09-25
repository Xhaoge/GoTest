package main

import "fmt"

func main(){
    fmt.Println("nima!")
    const LENGTH  int = 10 
    const WIDTH int = 5
    var area int
    const a,b,c =1,false,"str" // 多重赋值

    area = LENGTH * WIDTH
    fmt.Println(area)
    fmt.Println("面积为：%d",area)
    //Println("")
    fmt.Println(a,b,c)
}