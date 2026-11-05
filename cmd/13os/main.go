package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "demo_data"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "menu.txt")
	content := []byte("芋泥波波\n绿茶\n红茶\n")

	if err := os.WriteFile(path, content, 0644); err != nil {
		panic(err)
	}

	// 追加一行
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	_, _ = f.WriteString("爷爷不泡茶\n")
	f.Close()

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("文件内容:\n" + string(data))

	// 演示完清理
	_ = os.RemoveAll(dir)

	First()

	Second()

	// Third()

	Five()


}
