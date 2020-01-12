package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name 		string
	Age  		int
	Birthday 	string
	Sal  		float64
	Skill  		string
}

func testStruct(){
	monster := Monster{
		Name:"NIMA",
		Age:299,
		Birthday:"2011-11-11",
		Sal:800,
		Skill:"芭蕉扇",
	}
	// 将monster 序列化
	data ,err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误，err：",err)
	}
	// 输出序列化后的结果；
	fmt.Printf("monsrter 序列化后=%v\n",string(data))

}

func testMap(){
	//var a map[string]interface{}
	a := make(map[string]interface{})
	a["work"] = "程序员"
	a["address"] = "chengdu"
	data ,err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列化错误，err：",err)
	}
	// 输出序列化后的结果；
	fmt.Printf("map 序列化后=%v\n",string(data))
}

func testSlice(){
	slice := make([]map[string]interface{},0)
	m1 := make(map[string]interface{},1)
	m1["name"] = "xhagoe"
	m1["age"] = 10
	slice = append(slice,m1)
	slice = append(slice,m1)
	m2 := make(map[string]interface{},1)
	m2["name"] = "zhagnsan"
	m2["age"] = 12
	slice = append(slice,m2)
	data ,err := json.Marshal(&slice)
	if err != nil {
		fmt.Printf("序列化错误，err：",err)
	}
	// 输出序列化后的结果；
	fmt.Printf("slice 序列化后=%v\n",string(data),"长度为：",len(slice))
}

func main(){
	fmt.Println("json 序列化map test")
	testStruct()
	testMap()
	testSlice()
}