package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("========== 协程 ============")

	var wg sync.WaitGroup
	for _, file := range []string{"a.txt", "b.txt", "c.txt"} {
		wg.Go(func() {
			downLoadFile(file)
		})

	}
	wg.Wait()
}

func downLoadFile(file string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%s下载完成\n", file)
}
