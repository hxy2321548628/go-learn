package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Second() {
	dir := "tmp"
	if err := os.MkdirAll(dir, 0711); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "hello.txt")
	if f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0711); err != nil {
		panic(err)
	} else {
		defer f.Close()
		if _, err := f.WriteString("这是第一行\n"); err != nil {
			fmt.Println(err)
		}
		f.WriteString("这是第二行\n")
		f.WriteString("这是第三行\n")
	}
}
