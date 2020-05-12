package main

import (
	"fmt"
)

func quickSort(nums []int, left, right int) {
	val := nums[(left+right)/2]
	i, j := left, right
	for nums[j] > val {
		j--
	}
	for nums[i] < val {
		i++
	}
	nums[i], nums[j] = nums[j], nums[i]
	i++
	j--
	if i < right {
		quickSort(nums, i, right)
	}
	if j > left {
		quickSort(nums, left, j)
	}
}

func bubbleSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

func fibonacciRecuresion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacciRecuresion(n-1) + fibonacciRecuresion(n-2)
	}
}

func fibonacciSum(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fibonacciSum(n-1) + fibonacciSum(n-2)
	}
}

func main() {
	res := fibonacciRecuresion(3)
	fmt.Println(res)
	ressum := fibonacciSum(3)
	fmt.Println(ressum)
}
