package main

import "fmt"

func main() {
	fmt.Println("============== 数组和切片 ============")
	intArr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("intArr[4]: %v\n", intArr[3])
	fmt.Printf("len(intArr): %v\n", len(intArr))

	hobbit := []string{"唱", "跳", "rap"}
	fmt.Printf("hobbit: %v\n", hobbit)
	fmt.Printf("len(hobbit): %v\n", len(hobbit))
	fmt.Printf("cap(hobbit): %v\n", cap(hobbit))
	hobbit = append(hobbit, "篮球")
	fmt.Printf("hobbit: %v\n", hobbit)
	fmt.Printf("len(hobbit): %v\n", len(hobbit))
	fmt.Printf("cap(hobbit): %v\n", cap(hobbit))

	//数组满了之后会2倍2倍的扩容
	testSlice := []int{
		1, 2, 3, 4, 5,
	}
	testSlice = append(testSlice, 7)
	fmt.Printf("len(testSlice): %v\n", len(testSlice))
	fmt.Printf("cap(testSlice): %v\n", cap(testSlice))

	fmt.Println("============== 数组插入 ============")
	slice1 := []int{
		1, 2, 3, 4, 5,
	}
	slice1 = append([]int{-1, 0}, append(slice1[:3], append([]int{100, 200}, slice1[3:]...)...)...)
	fmt.Printf("slice1: %v\n", slice1)

	score := []int{95, 58, 76, 42, 88}
	num := 0
	for index, s := range score {
		if s < 60 {
			score = append(append(score[:index], 0), score[index+1:]...)
		} else {
			num += 1
		}
	}
	fmt.Printf("score: %v\n", score)
	fmt.Println("有",num,"门合格")

	fmt.Println("============== 传参区别 ============")
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println("调用后数组:", arr) // [1 2 3] 不受影响

	sl := []int{1, 2, 3}
	modifySlice(sl)
	fmt.Println("调用后切片:", sl) // [999 2 3] 受影响

}
func modifyArray(a [3]int) { // 如果改成 指针类型，则传参就是饮用，而不是copy
	a[0] = 999
	fmt.Println("函数内数组:", a)
}

func modifySlice(s []int) {
	s[0] = 999
	fmt.Println("函数内切片:", s)
}
