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

func get_average(arr []int, size int) float32 {
	var sum int
	var avg float32
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}

func fibonacci(n int) int {
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	var arr = [5]int{1, 2, 3, 4, 5}

	ret = max(a, b)
	fmt.Printf("max = %d\n", ret)

	// 函数传入数组
	res := get_average(arr[:], 5)
	fmt.Printf("average = %.2f\n", res)
	n := 3
	fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
}
