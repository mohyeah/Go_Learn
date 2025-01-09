package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 声明变量并初始化
	var a string = "Hello"
	fmt.Println(a)

	// 声明变量，未初始化, 默认值为0
	var b, c int
	fmt.Println(b, c)
	// bool ,默认为False, 即0
	var d bool
	fmt.Println(d)

	var i int
	var f float64
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, d, s)

	intVal := 1
	fmt.Println(intVal)

	// 多变量声明, 该因式分解关键字的写法一般用于函数外声明全局变量
	// a, b := 12, 12 该格式仅能在函数体中出现
	var (
		a1 int = 1
		a2 int = 2
	)
	fmt.Println(a1, a2)

	var h, k int
	var j string
	h, k, j = 1, 2, "hello"
	fmt.Println(h, k, j)

	// 交换h, k的值
	h, k = k, h

	// 常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const m, n, q = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	fmt.Println()
	fmt.Println(m, n, q)

	// 常量可用作枚举
	// iota从0开始递增
	const (
		apple = iota
		orange
		peach
	)
	fmt.Println(apple, orange, peach)

	const (
		rho     = iota
		theta   = "ha" // 该值会覆盖掉之前的iota, iota++
		epsilon        // iota++
		ita     = 100  // 覆盖掉之前的iota, iota++
		lambda         // iota++
		mu      = iota //恢复iota计数
	)
	fmt.Println(rho, theta, epsilon, ita, lambda, mu)

	const (
		alpha   = "abc"
		beta    = len(alpha)
		charlie = unsafe.Sizeof(alpha)
	)
	fmt.Println(apple, beta, c)

	// x << n => x*(2^n)
	const (
		ek1 = 1 << iota // ek1=1, 左移0位, 不变
		ek2 = 3 << iota // ek2=3, 左移1位, 变为110 -> 6
		ek3             // ek3=3, 左移2位, 变为1100 -> 12
		ek4             // ek4=3, 左移3位, 变为11000 -> 24
	)

	// 运算符
	// << 左移运算符: x << n => x * (2^n)

	//指针
	// & 取地址
	// * 取值
	ak := 1
	ptr := &ak
	fmt.Printf("ak的值为:%d\n", ak)
	fmt.Printf("ptr为:%d\n", ptr)
	fmt.Printf("*ptr为:%d\n", *ptr)

	// 条件语句: if, switch, select(通信选择,随机执行一个可运行的case,若无,则阻塞)
	// 循环语句: for, range, while, do while
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

}
