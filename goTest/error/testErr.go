package main

import ("fmt"
		"errors")

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
//函数去读取以配置文件init.conf 的信息；
//如果传入的文件名不正确，我们就返回一个自定义的错误；

func readConf(name string)(err error){
	if name == "config.ini"{
		//读取；
		return nil
	}else{
		//返回一个自定义错误；
		return errors.New("读取文件错误。。。。")
	}
}

func test02(){
	err := readConf("config2.ini")
	if err != nil{
		//读取文件发送错误，就输出这个错误，并终止程序；
		panic(err)
	}
	fmt.Println("test02() continue")
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
	test02()
}
