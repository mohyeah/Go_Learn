package calutils

import (
	"errors"
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

func Div1(n1 int, n2 int) (result int) {
	// 利用defer和recover函数,捕获异常 利用 defer 匿名函数 来调用
	defer func() {
		// 调用recover函数,捕获异常
		err := recover()
		// 如果捕获到异常,则返回错误信息,否则返回0
		if err != nil {
			err = fmt.Errorf("division by zeror")
			//fmt.Println("error is:", err)
			// 立即终止程序, 抛出异常
			panic(err)
		}
	}()
	result = n1 / n2
	return result
}
func Div2(n1 int, n2 int) (result int, err error) {
	if n2 == 0 {
		// 自定义错误
		return 0, errors.New("division by zeror")
	} else {
		result = n1 / n2
		return result, nil
	}
}
