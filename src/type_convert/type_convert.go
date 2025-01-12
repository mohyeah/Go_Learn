package main

import (
	"fmt"
	"reflect"
)

// 定义接口 Writer
type Writer interface {
	Write([]byte) (int, error)
}

// 实现 Writer 接口的 StringWriter 结构体
type StringWriter struct {
	str string
}

// 实现 Writer 接口的 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

// 定义接口
type Speaker interface {
	Speak()
}

// 定义结构体
type Person struct {
	Name string
}

// Person 实现了 Speaker 接口的 Speak 方法
func (p Person) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	type myInt int
	var i myInt = 10
	var j = 10
	// 自定义类型 myInt 和 int 不同的类型
	fmt.Println(reflect.TypeOf(i) == reflect.TypeOf(j))

	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello, world"
	fmt.Println(sw.str)

	p1 := Person{Name: "SpongeBob"}
	var speaker1 Speaker = p1
	speaker1.Speak()
}
