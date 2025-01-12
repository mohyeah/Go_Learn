package main

import (
	"fmt"
	"reflect"
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

func scoreClassify(score int) string {
	switch score / 10 {
	case 10:
		return "A"
	case 9:
		return "A"
	case 8:
		return "B"
	case 7:
		return "C"
	case 6:
		return "D"
	default:
		return "E"
	}
}

// 多元参数
func add(args ...int) {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Println(sum)
}

// 函数也是一种类型
func test(n int) {
	fmt.Println(n)
}

// 嵌套
func test_2(n1 int, n2 float64, testFunc func(int)) {
	fmt.Println("test_2", n1)
}

// 函数返回值命名
func product(n1 int, n2 int) (result int) {
	result = n1 * n2
	return result
}

func main() {
	var a int = 100
	var b int = 200
	var ret int
	fmt.Println(swap("a", "b"))

	var arr = [5]int{1, 2, 3, 4, 5}

	ret = max(a, b)
	fmt.Printf("max = %d\n", ret)

	// 函数传入数组
	res := get_average(arr[:], 5)
	fmt.Printf("average = %.2f\n", res)
	n := 3
	fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))

	score := 89
	fmt.Printf("score = %d, classify = %s\n", score, scoreClassify(score))

	add(1, 2, 3, 4, 5)

	t1 := test
	fmt.Println("type of t1 is:", reflect.TypeOf(t1))
	t1(10)
	test_2(1, 4.5, t1)

}
