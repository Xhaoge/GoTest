package main

import (
	"fmt"
	"reflect"
)
type SearchCase struct{

}


type SearchResponse struct {
	Msg	string
	routing []*routing
}

type routing struct {
	FromCity 	string
	ToCity		string
	StartDate   string
	RetDate		string
}

func main() {
	var SearchCase SearchCase
	fmt.Println(SearchCase)
	cc := reflect.ValueOf(&SearchCase)
	ss := reflect.TypeOf(SearchCase).String()
	fmt.Println(cc)
	fmt.Println(ss) // 返回*main.SearchCase

	dd := reflect.TypeOf(&SearchCase).Elem().Name()
	fmt.Println(dd) // 返回SearchCase 
	aa := reflect.TypeOf(dd)
	fmt.Println(aa) // string

	
	rout := routing{"BJs","lax","20191225","20191229"}
	sr := SearchResponse{}
	fmt.Println("sr:",sr)
	fmt.Println("rout:",rout)
	
}

