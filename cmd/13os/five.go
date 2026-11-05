package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Five() {
	dir := "tmp/a/b/c"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(path, []byte("这是一个测试文件\n"), 0755); err != nil {
		panic(err)
	}

	//遍历
	if dirs, err := os.ReadDir(dir); err != nil {
		panic(err)
	} else {
		for _, dir := range dirs {
			fmt.Printf("dir: %v\n", dir)
		}
	}

	//删除
	if err := os.RemoveAll("tmp/a"); err != nil {
		panic(err) 
	}
	

}
