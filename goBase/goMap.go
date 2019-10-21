package main

import "fmt"

// map 类型test

func printM(cc map[string]string) {
	for c := range cc {
		fmt.Println("capital of ", c, "是", cc[c])
	}
}

func main() {
	fmt.Println("map 学习；")
	// 申名变量；
	var countryCapitalMap map[string]string
	// 创建集合
	countryCapitalMap = make(map[string]string)
	//map charu key-key键值对，各个国家对应的首都；
	countryCapitalMap["France"] = "Pairs"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Tokyo"] = "Tokyo"
	countryCapitalMap["japan"] = "New Delhi"
	// 使用key 输出map值；
	for country := range countryCapitalMap {
		fmt.Println("capital of ", country, "is", countryCapitalMap[country])
	}
	// 查看元素在集合中是否存在；
	captial, ok := countryCapitalMap["China"]
	if ok {
		fmt.Println("capital fo china is ", captial)
	} else {
		fmt.Println("capital of China is not present")
	}
	countryCapitalMap["yayaya"] = "lele"
	fmt.Println(countryCapitalMap)

	// 删除元素
	delete(countryCapitalMap, "yayaya")
	printM(countryCapitalMap)
	fmt.Printf("map 的长度为：%d\n", len(countryCapitalMap))
	countryCapitalMap["yaya"] = "le"
	fmt.Printf("map 的长度为：%d\n", len(countryCapitalMap))
}
