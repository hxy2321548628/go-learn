package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("============= 通道 ============")
	// deadLock()
	corrChan()

	n := 2
	ch1 := make(chan int)
	go func(n int) {
		ch1 <- n * 2
	}(n)
	n = <-ch1
	fmt.Printf("n: %v\n", n)

	go func() {
		fmt.Println("nihao")
		// 这一行注释就会死锁，因为通道是无缓冲的
		fmt.Printf("(<-ch1): %v\n", (<-ch1))
	}()
	ch1 <- n

	ch2 := make(chan int, 5)
	for v := range 5 {
		ch2 <- v
	}
	close(ch2) // 关键：关闭后 range 才能正常退出
	for v := range ch2 {
		fmt.Printf("v: %v\n", v)
	}

	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch4 <- 4
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch3 <- 3
	}()

	for range 2 {
		select {
		case n := <-ch4:
			fmt.Println("ch2:", n)
		case n := <-ch3:
			fmt.Println("ch3:", n)
		}
	}

}

func corrChan() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Go(func() {
		ch <- 1
	})

	// 先获取
	res := <-ch
	wg.Wait()

	fmt.Println(res)
}
