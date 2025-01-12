package calutils

import (
	"fmt"
)

// 大写函数可以被其他包调用
func Product(n1 int, n2 int) (result int) {
	result = n1 * n2
	return result
}

func Sum(n1 int, n2 int) (result int) {
	result = n1 + n2
	return result
}

func Sub(n1 int, n2 int) (result int) {
	result = n1 - n2
	return result
}

func Div(n1 int, n2 int) (result int, err error) {
	if n2 == 0 {
		return 0, fmt.Errorf("division by zeror")
	}
	result = n1 / n2
	return result, nil
}
