package main

import (
	"fmt"
)

/*
先从数列中取出一个数作为基准数。
分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
再对左右区间重复第二步，直到各区间只有一个数。*/

func quickSort(arr []int, right, left int) {
	if right < left {
		i, j := right, left
		key := arr[(right+left)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if right < j {
			quickSort(arr, right, j)
		}
		if i < left {
			quickSort(arr, i, left)
		}
	}
}

// 切片排序，从大到小
func quickSortFromMaxToMin(arr []int, left, right int) {
	if left < right {
		i, j := left, right
		key := arr[(right+left)/2]
		for i <= j {
			for arr[i] > key {
				i++
			}
			for arr[j] < key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if left <= j {
			quickSortFromMaxToMin(arr, left, j)
		}
		if right >= i {
			quickSortFromMaxToMin(arr, i, right)
		}
	}
}

func main() {
	fmt.Println("quick sort......")
	arr := []int{3, 7, 9, 8, 38, 93, 12, 3, 222, 222, 45, 93, 23, 84, 65, 2}
	fmt.Println("len: ", len(arr))
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("排序后结果： ", arr)
	quickSortFromMaxToMin(arr, 0, len(arr)-1)
	fmt.Println("排序后结果： ", arr)
	fmt.Println(2 > "abcd")
}
