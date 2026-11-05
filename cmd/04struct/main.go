package main

import (
	"fmt"
)

func main() {
	fmt.Println("=============== 结构体 ==============")

	chinese := Book{"语文", "hhh", 20}
	english := Book{name: "英语", author: "yyy", pages: 21}
	fmt.Printf("chinese: %v\n", chinese)
	fmt.Printf("english: %v\n", english)

	hxy := User{
		name: "hxy",
		address: Address{
			city:  "广州市",
			block: "白云区",
		},
	}
	fmt.Printf("%s 住在 %s%s\n", hxy.name, hxy.address.city, hxy.address.block)

	chinese.Info()
	chinese.AddPage(1)
	chinese.Info()

	r := NewRetangle(WithLong(2), WithWidth(3))
	fmt.Printf("r.Area(): %v\n", r.Area())
	fmt.Printf("r.Perimeter(): %v\n", r.Perimeter())
	fmt.Printf("r: %v\n", r)
}

type Rectangle struct {
	long  float64
	width float64
}

func NewRetangle(opts ...Option) *Rectangle {
	// 1. 先设置默认值
	r := &Rectangle{
		long:  2,
		width: 1,
	}

	// 2. 依次应用每个 Option
	for _, opt := range opts {
		opt(r)
	}

	// 3. 用最终配置构造 User
	return r
}

type Option func(*Rectangle)

func WithLong(long float64) Option {
	return func(r *Rectangle) { r.long = long }
}
func WithWidth(width float64) Option {
	return func(r *Rectangle) { r.width = width }
}

func (r *Rectangle) Area() float64 {
	return r.long * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.long + r.width)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{ long = %v, width = %v }", r.long, r.width)
}

type User struct {
	name    string
	address Address
}

type Address struct {
	city  string
	block string
}

type Book struct {
	name   string
	author string
	pages  int
}

func (b Book) Info() {
	fmt.Println(b)
}

func (b *Book) AddPage(num int) {
	b.pages += 1
}

func (b Book) String() string {
	return fmt.Sprintf("Book{ name=%v, author=%v, pages=%v}\n", b.name, b.name, b.pages)
}
