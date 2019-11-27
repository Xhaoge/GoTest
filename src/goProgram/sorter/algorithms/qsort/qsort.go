package main

import (
	"flag"
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i","infile","File contains values for sorting")
var outfile * string = flag.String("o","outfile","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")

// 读取文件中的数据
func readValues(infile string)(values []int,err error){
	file, err := os.Open(infile)
	if err !=nil {
		fmt.Println("Failed to open the input file ",infile)
		return
	}

	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int,0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if errl != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seems unexpected.")
			return
		}
		str :+ string(line) //转换字符数组为字符串
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values,value)
	}

}

//输出数据
func writeValues(values []int,outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed to create the output file",outfile)
		return err
	}

	defer file.Close()

	for _,value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}



func main(){
	flag.Parse()

	if infile != nil{
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil{
		fmt.Println("readValues:",values)
	}else {
		fmt.Println("err:",err)
	}
}

