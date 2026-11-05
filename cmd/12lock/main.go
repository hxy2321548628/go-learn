package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(arr *[]string, wg *sync.WaitGroup, mu *sync.Mutex, tea string) {
	defer wg.Done()
	mu.Lock()                // 上锁：别人先别动 arr
	*arr = append(*arr, tea) // 临界区：只有拿到锁的协程能改
	mu.Unlock()              // 解锁：下一位可以进来
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	arr := []string{"芋泥波波", "绿茶", "红茶"}

	wg.Add(3)
	go worker(&arr, &wg, &mu, "柠檬茶")
	go worker(&arr, &wg, &mu, "霸王茶姬")
	go worker(&arr, &wg, &mu, "爷爷不泡茶")
	wg.Wait()

	fmt.Println(arr)

	var lock sync.RWMutex
	myMap := map[string]int{
		"hxy": 1,
		"hhh": 4,
	}
	wg.Go(func() {
		Read(myMap, &lock)
	})
	wg.Go(func() {
		Read(myMap, &lock)
	})
	wg.Go(func() {
		Write(myMap, "hxy", &lock)
	})
	wg.Go(func() {
		if err := Write(myMap, "hxh", &lock); err != nil {
			fmt.Println("error: ", err)
		}
	})

	wg.Wait()

	fmt.Println(myMap)

}

func Read(dic map[string]int, lock *sync.RWMutex) {
	time.Sleep(20 * time.Millisecond)
	lock.RLock()
	defer lock.RUnlock()
	for key, value := range dic {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
	fmt.Println("-----------------------")
}

func Write(dic map[string]int, key string, lock *sync.RWMutex) error {
	time.Sleep(200 * time.Millisecond)
	lock.Lock()
	defer lock.Unlock()
	if value, ok := dic[key]; !ok {
		return fmt.Errorf("%s不存在", key)
	} else {
		dic[key] = value + 1
	}
	return nil
}
