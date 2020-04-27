package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func TimeFormatStr(t1 timestamp.Timestamp, t2 string){
	fmt.Println("t1: ",t1)
	fmt.Println("t2: ",t2)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-11", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())
}

func main(){
	tstamp := time.Now().Unix()
	TimeFormatStr(tstamp,"ssss")
}