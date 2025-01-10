package main

import "fmt"

func main() {
	var a int = 10
	var pa *int = &a

	fmt.Printf("a = %d\n", a)
	fmt.Printf("pa = %d\n", pa)
	fmt.Printf("*pa = %d\n", *pa)

	// 未初始化指针, 其值为nil
	var ptr *int
	fmt.Printf("ptr = %d\n", ptr)
	fmt.Printf("%t", ptr == nil)
}
