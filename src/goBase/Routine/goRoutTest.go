package main

import (
	"fmt"
	"time"
)


var Arr1 = []string{"xhaoge","nima","golang"}
type Xhaoge struct{
	name  string
	age   int
}

typeFlower struct{
	[]Xhaoge
}

func (w *Flower) find(){
	fmt.Println("@@@",w[0])
	fmt.Println("@@@",w[1])
	fmt.Println("@@@",w[2])
	for _,n := range w{
		go func() {
			fmt.Println("go func() Println",n)
		}()
		fmt.Println("for Println",n)
	}
}


func main(){
	fmt.Println("test for gorouting")
	fmt.Println(Arr1)
	var ww = []Xhaoge{}
	var w1 = Xhaoge{"chanshi",18}
	var w2 = Xhaoge{"kobi",17}
	var w3 = Xhaoge{"manba",16}
	ww = append(ww,w1,w2,w3)
	fmt.Println("ww:",ww)
	for _, m := range Arr1{
		go func() {
			fmt.Println("go func() Println",m)
		}()
		go func() {
			fmt.Println("go func()1 Println",m)
		}()
		time.Sleep(time.Second)
		fmt.Println("for Println",m)
	}

}