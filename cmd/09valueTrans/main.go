package main

import (
	"fmt"
)

func main() {
	fmt.Println("=========== 值传递和引用传递 ============")

	s := make([]int, 0, 4) // len=0, cap=4
	appendOne(s)
	fmt.Println(s, len(s)) // [] 0 —— 注意：形参 s 的 len 变了，但 main 里的 s 仍是原来的 header
	fmt.Println(s[:1])     // [100]  ← 手动把 len 撑到 1，就能看到数据了

	scoreTable := map[string][]int{
		"hxy": {12, 13, 14},
	}
	addScore(scoreTable, "hxy", 111)
	fixFirst(scoreTable, "hxy", 100)
	fmt.Printf("scoreTable: %v\n", scoreTable)
	addScore(scoreTable, "hxy", 111)
	addScore(scoreTable, "hxy", 111)
	addScore(scoreTable, "hhh", 111)
	fmt.Printf("scoreTable: %v\n", scoreTable)

}

func addScore(scoresTable map[string][]int, name string, score int) error {
	scores, ok := scoresTable[name]
	if !ok {
		return fmt.Errorf("%s 不存在", name)
	}
	// 这里一定要赋值， 不然就只是修改了局部变量，并没有修改原来的map
	scoresTable[name] = append(scores, score)
	return nil
}

func fixFirst(scoresTable map[string][]int, name string, newScore int) error {
	scores, ok := scoresTable[name]
	if !ok {
		return fmt.Errorf("%s 不存在", name)
	}
	scores[0] = newScore
	return nil
}

func appendOne(s []int) {
	s = append(s, 100) // 在共享的底层数组上追加
}
