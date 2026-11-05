package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	fmt.Println("========== 函数 ==========")

	for i := range 3 {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

	sum := 0
	var lock sync.Mutex
	var wg sync.WaitGroup
	for range 100 {
		wg.Go(func() {
			lock.Lock() // 加锁， 并发安全
			defer lock.Unlock()
			sum++
		})
	}
	wg.Wait()
	fmt.Printf("sum: %v\n", sum)

	fmt.Println("========== 闭包 ==========")
	counter := makeCounter()
	fmt.Printf("counter(): %v\n", counter())
	fmt.Printf("counter(): %v\n", counter())

	a, b := 1, 2
	fmt.Printf("max(a, b): %v\n", myMax(a, b))

	c, d := 1, 2
	c, d = swap(c, d)
	fmt.Println("swap(c,d): ", c, d)

	var n float64 = -1
	if _, err := sqrt(n); err != nil {
		fmt.Println(err)
	}

}

func sqrt(n float64) (float64, error) {
	if n <= 0 {
		return 0, errors.New("不能为0")
	} else {
		return math.Sqrt(n), nil
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func myMax(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func makeCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
