package main

import "fmt"
/*条件语句test*/

const(
	a = true
	b = false
)

func main(){
	if (a&&b){
		fmt.Println("第一行-条件为：")
	}
	if (a||b){
		fmt.Println("第二行")
	}
	if (a&&b){
		fmt.Println("nima")
	}else{
		fmt.Println("baibia")
	}
	if (!(a&&b)){
		fmt.Println("!不需要；；；；")
	}



}