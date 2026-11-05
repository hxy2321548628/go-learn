package main

import (
	"fmt"
)

func main() {

	fmt.Println("============== 变量定义 =============")
	var name string = "hxy"
	fmt.Println("姓名： ", name)

	age := 23
	fmt.Println("年龄为：", age)

	school, major := "中南财经政法大学", "金融工程"
	fmt.Printf("就读于%s, 专业为%s \n", school, major)

	var tint = 123
	var tst = &tint
	fmt.Println(tst)

	fmt.Println("============== 数组切片 =============")

	//数组
	arr := [5]int{1, 2, 5}
	fmt.Println("数组: ", arr)
	fmt.Printf("cap(arr): %v\n", cap(arr))
	fmt.Printf("len(arr): %v\n", len(arr))

	//切片
	var arr2 []int
	arr2 = append(arr2, 1, 2, 3)
	//切片会自动扩容，数组的容量和大小是固定的
	fmt.Printf("cap(arr2): %v\n", cap(arr2))
	fmt.Printf("len(arr): %v\n", len(arr))

	fmt.Println("============== 字典映射 =============")

	//字典 映射
	var myMap map[int]int = make(map[int]int)
	myMap[1] = 1
	myMap[2] = 1
	myMap[3] = 1
	myMap[4] = 1
	myMap[5] = 1
	myMap[6] = 1
	fmt.Printf("myMap: %v\n", myMap)
	fmt.Printf("len(myMap): %v\n", len(myMap))

	//第二种方式
	myMap2 := map[int]int{
		1: 2,
		2: 3,
		3: 4,
	}
	fmt.Printf("myMap2: %v\n", myMap2)
	//获取值 获取不存在的值时，不会报错，返回0值，通过err来判断是否存在
	var value, err = myMap2[4]
	fmt.Printf("value: %v\n", value)
	fmt.Printf("err: %v\n", err)
	//删除操作
	delete(myMap2, 3)
	fmt.Printf("myMap2: %v\n", myMap2)
	for k := range myMap2 {
		delete(myMap2, k)
	}
	fmt.Printf("myMap2: %v\n", myMap2)

	fmt.Println("============== 结构体 =============")

	var hxy = &Person{name: "hxy", age: 23}
	fmt.Printf("hxy: %v\n", hxy)
	hxy.name = "adsda" // 可以通过分包进行隔离，go中只有大写开头的才能被外部进行访问
	fmt.Printf("hxy: %v\n", hxy)

	fmt.Println("============== 接口 =============")
	dog := Dog{}
	fmt.Printf("dog.Speak(): %v\n", dog.Speak())

	bookList := [3]*Book{ //注意这里时指针类型
		{name: "Chinese"},
		{name: "English"},
		{name: "Math"},
	}
	fmt.Printf("bookList: %v\n", bookList)
	for _, book := range bookList {
		fmt.Println(book)
	}
}

type Book struct {
	name string
}

func (b *Book) String() string {
	return fmt.Sprintf("Book{name: %q}", b.name)
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪"
}

type Person struct {
	name string
	age  int
}

// 实现 fmt.Stringer 接口
func (p *Person) String() string {
	return fmt.Sprintf("Person{name: %q, age: %d}", p.name, p.age)
}
