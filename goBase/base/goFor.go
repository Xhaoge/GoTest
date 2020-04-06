package main 

import "fmt"
/* for 语句 test*/


func main(){
	//var a int8;
	numbers :=[6]int{1,2,3,7}
	for a:=1;a<=10;a++{
		fmt.Printf("计数a = %d\n",a)
		for i,x :=range numbers{
			fmt.Printf("第%d位的值为：%d,和为：%d\n",i,x,x+a)
			if a+x==11{
				break
			}
		}
	}

	
	


}