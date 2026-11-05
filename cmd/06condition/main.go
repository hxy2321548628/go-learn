package main

import "fmt"

func main() {
	fmt.Println("================== 条件分支 ===============")
	for _, score := range []int{50, 70, 88, 96} {
		Check(score)
	}

}

func Check(score int) {
	fmt.Println(score, ":")
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	fmt.Println("")
}
