package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Name string `json:name`
}

func main() {
	js := `{
		"Name":"xhoage"
	}`
	var p people
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("people: ", p)
}
