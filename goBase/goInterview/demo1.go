package main

import "fmt"

type param map[string]interface{}

type show struct {
	param
}

type student struct {
	name string
}

func testStudent(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Println(msg.name)
	}
}

func main() {
	s := new(show)
	s.param = make(map[string]interface{})
	s.param["xhaoge"] = 100
	fmt.Println(s.param)
	myS := student{
		name: "xhaoge",
	}
	testStudent(myS)
}
