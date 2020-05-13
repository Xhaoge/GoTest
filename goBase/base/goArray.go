package main

import "fmt"

//数组 test 需要指明数组的大小和存储的数据类型

type nnn struct {
	Name string
}

func main() {
	var balance [10]float32
	fmt.Printf("balance = %.1f\n", balance)
	var balanc = [5]float32{100.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("balance1 = %.2f\n", balanc)
	balanc[3] = 2
	fmt.Printf("balance2 = %.2f\n", balanc)
	fmt.Printf("balance 的第三位是：%.2f\n", balanc[2])

	var i int
	//根据for下标遍历数组
	for i = 0; i < len(balanc); i++ {
		fmt.Printf("%d for balance: %.2f\n", i, balanc[i])
	}
	//使用range 函数遍历数组值；
	for _, v := range balanc {
		fmt.Printf("依次打印数组中的值：%f", v)
	}
	a := balanc
	fmt.Println("\n", a)

	// 切片，slice;
	darr := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before: ", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after: ", darr, "切片容量为：", cap(darr), "切片长度为：", len(darr))
	fmt.Println("old slice : ", dslice)
	dslice = append(dslice, 23, 54, 56, 8888, 7, 87, 99, 678, 50, 29, 66, 109)
	fmt.Println("new slice : ", dslice, "切片容量为：", cap(dslice), "切片长度为：", len(dslice))

	var arrS []interface{}
	var v nnn
	v.Name = "xhaoge"
	arrS = append(arrS, v)
	arrS[0] = v
	fmt.Println("arrS:", arrS)

	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	fmt.Println(len(Slice_a), cap(Slice_a))

	// array_b := [3]int{1, 2, 3}
	// array_b = append(array_b, 44)
	// fmt.Println("array_b: ", array_b)
}
