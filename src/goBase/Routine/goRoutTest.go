package main

import (
	"fmt"
	"time"
)


var Arr1 = []string{"xhaoge","nima","golang"}

func main(){
	fmt.Println("test for gorouting")
	fmt.Println(Arr1)
	for _, m := range Arr1{
		go func() {
			fmt.Println("go func() Println",m)
		}()
		time.Sleep(time.Second)
		fmt.Println("for Println")
	}

}