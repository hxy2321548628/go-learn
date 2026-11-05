package main

import "fmt"

func main() {
	fmt.Println("============= 异常处理 ============")
	mustPositive(-1)

	fmt.Println("继续运行 main 函数")
}

func mustPositive(n int) {
	defer func() {
		if info := recover(); info != nil {
			fmt.Printf("错误被捕获: %v\n", info)
		}
	}()

	if n < 0 {
		panic("必须为整数")
	}

	fmt.Println("这是后续代码，如果发生了 panic 就不会运行")
}
