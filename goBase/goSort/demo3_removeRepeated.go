package main

import (
	"fmt"
	"sort"
)

func removeRepeated(arr []string) []string {
	newArr := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i+1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 去除重复字段和空字符串；先排序，再去重
func removeRepeatedBySort(arr []string)[]string{
	newArr := make([]string,0)
	sort.Strings(arr)
	for i:=0;i<len(arr);i++{
		if (i>0 && arr[i]==arr[i-1]) || len(arr[i])==0{
			continue
		}
		newArr = append(newArr,arr[i])
	}
	return newArr
}

// 通过map 主键唯一的特性过滤重复元素
func removeRepeatedByMap(arr []string)[]string{
	newArr := make([]string,0)
	temMap := make(map[string]interface{})
	for _,val := range arr{
		// 判断主键是否存在
		if _, ok := temMap[val]; !ok {
			newArr = append(newArr, val)
			temMap[val] = nil
		}
	}
	return newArr
}

func main() {
	fmt.Println("remove repeated number")
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	fmt.Println(removeRepeated(arr))
	arr2 := []string{"hello", "", "world", "yes", "hello", "nihao", "shijie", "hello", "yes", "nihao","good"}
	fmt.Println(removeRepeatedBySort(arr2))
	var arr3 = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	fmt.Println(removeRepeatedByMap(arr3))
}
