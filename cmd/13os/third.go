package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Third() {
	dir := "tmp"
	if err := os.MkdirAll(dir, 0711); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "third.txt")
	if _, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0711); err != nil {
		panic(err)
	} else {
		fmt.Println("第一次创建成果")
	}

	if _, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0711); err != nil {
		fmt.Println("第二次创建失败")
		panic(err)
	} else {
		fmt.Println("第二次创建成果")
	}

}
