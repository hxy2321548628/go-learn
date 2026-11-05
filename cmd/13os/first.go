package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func First() {
	dir := "tmp"
	if err := os.MkdirAll(dir, 0711); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "hello.txt")
	content := []byte("hello golang\n")

	if err := os.WriteFile(path, content, 0711); err != nil {
		panic(err)
	}

	if data, err := os.ReadFile(path); err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}

}
