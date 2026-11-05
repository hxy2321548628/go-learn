package main

import (
	"fmt"
	"reflect"
)

func main() {

	var age Age = 1
	t := reflect.TypeOf(age)
	fmt.Printf("t.Name(): %v\n", t.Name())
	fmt.Printf("t.Kind(): %v\n", t.Kind())

	if t.Kind() == reflect.Int {
		fmt.Println("一定是一个 int 类型的值")
	}

	// 反射修改值
	v := reflect.ValueOf(&age).Elem()
	if v.CanSet() && v.IsValid() && v.Kind() == reflect.Int {
		v.SetInt(10086)
	}
	fmt.Printf("age: %v\n", age)

	nv := v.Interface().(Age)
	fmt.Printf("nv: %v\n", nv)

	//通过反射调用方法
	call := v.MethodByName("Say")
	if call.IsValid() {
		call.Call([]reflect.Value{
			reflect.ValueOf("你好"),
			reflect.ValueOf(1),
		})
	}

}

type Age int

func (a Age) Say(params string, n int) {
	fmt.Println(params, n)
}
