package main

import (
	"fmt"
	"strconv"
)

func max(n1, n2 int) int {
	// 声明局部变量
	var result int
	if n1 > n2 {
		result = n1
	} else {
		result = n2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func is_palindorme(x int) bool {
	var bytes []byte = []byte(strconv.Itoa(x))
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("max = %d\n", ret)
}
