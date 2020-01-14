package main

import (
	"fmt"
	"reflect"
)

type person struct{
	Name    string
}

func (p person)test(){
	fmt.Println("test struct")
}



func main(){
	fmt.Println("test")
	m1 := person{"xhaoge"}
	m2 := person{"sunhao"}
	var searchList []interface{}
	searchList =  append(searchList,m1,m2)
	Caselist :=make(map[string]interface{})
	fmt.Println(searchList)
	for _,g := range searchList{
		getType := reflect.TypeOf(g)
		getValue := reflect.ValueOf(g)
		fmt.Println(reflect.TypeOf(getType))
		knd := getType.Kind()
		fmt.Println("knd:",knd)
		for i:=0;i<getType.NumField();i++{
			field := getType.Field(i)
			value := getValue.Field(i).Interface()
			fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)

		}
		//获取方法 1.先获取interface的reflect type 然后通NumMethod进行遍历
		for i:=0;i<getType.NumMethod(); i++{
			m := getType.Method(i)
			fmt.Printf("%s: %v\n", m.Name, m.Type)
		}
	}
	
	Caselist["search"] = searchList

	fmt.Println("Caselist",Caselist)
	t := reflect.TypeOf(Caselist["search"]).Kind()
	s := reflect.ValueOf(Caselist["search"])
	fmt.Println(t)
	fmt.Println(s,reflect.TypeOf(s))
	// for _,n := range s{
	// 	c := reflect.TypeOf(n)
	// 	fmt.Println(c)
	// }

}