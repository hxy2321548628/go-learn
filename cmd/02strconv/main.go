package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("============= 类型转换 ============")

	age := 23
	ageString := strconv.Itoa(age)
	fmt.Println("int to string: ")
	fmt.Printf("strconv.Itoa(age): %T, %v\n", ageString, ageString)

	if _, err := strconv.ParseBool("111"); err != nil {
		fmt.Printf("bool值解析失败: %v\n", err)
	}

	var tfloat float64 = 9.99
	tint := int(tfloat)
	tfloat = float64(tint)
	fmt.Printf("tfloat: %v\n", tfloat)

	price, discount := "9.99", "true"
	parsePrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return
	}
	parseDiscount, err := strconv.ParseBool(discount)
	if err != nil {
		return
	}
	good := &Good{
		discountPrice: strconv.FormatFloat(parsePrice*0.8, 'f',6, 64), //单引号时字符，双引号才是字符串
		discount:      parseDiscount,
	}
	fmt.Printf("good: %v\n", good)

}

type Good struct {
	discountPrice string
	discount      bool
}

func (g *Good) String() string {
	return fmt.Sprintf("Good{discountPrice:%s, discount:%v}\n", g.discountPrice, g.discount)
}
