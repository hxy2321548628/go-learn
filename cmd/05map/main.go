package main

import (
	"fmt"
)

func main() {
	fmt.Println("============= 映射 ===========")
	var myMap1 map[string]int // 只声明， 并没有初始化， 不能存值
	fmt.Printf("myMap1: %v\n", myMap1)

	// myMap1["a"] = 1.
	myMap1 = make(map[string]int, 10)
	myMap1["a"] = 1
	fmt.Printf("myMap1: %v\n", myMap1)
	fmt.Println(len(myMap1))

	info := map[string]string{
		"name":  "hxy",
		"city":  "guangzhou",
		"hobit": "run",
	}
	info["city"] = "shenzhen"
	info["age"] = "23"
	fmt.Printf("info: %v\n", info)

	scores := map[string]int{
		"语文": 90,
		"英语": 88,
	}
	for _, subject := range []string{"语文", "数学", "英语"} {
		if score, ok := scores[subject]; !ok {
			fmt.Printf("%s不存在\n", subject)
		} else {
			fmt.Printf("%s得分%d\n", subject, score)
		}

	}

	fruitCount := make(map[string]int)
	for _, fruit := range []string{"苹果", "香蕉", "苹果", "橘子", "香蕉", "苹果"} {
		if _, ok := fruitCount[fruit]; !ok {
			fruitCount[fruit] = 1
		} else {
			fruitCount[fruit] += 1
		}
	}
	fmt.Printf("fruitCount: %v\n", fruitCount)

}
