package main

import "fmt"

// 定义一个 DivideError 结构；
type DivideError struct {
	dividee int
	divider int
}

// 实现‘error’接口
func (de *DivideError) Error() string {
	strFormat := "cannot proceed,the divider is zero dividee:%d devider:0"
	return fmt.Sprintf(strFormat, de.dividee)
}

//define "int" 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	fmt.Println("nima")
	// 正常情况下
	if result, err := Divide(100, 10); err == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, err := Divide(100, 0); err != "" {
		fmt.Println("errorMsg is :", err)
	}
}
