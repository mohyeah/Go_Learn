package main

import (
	"Go/src/base/pkg/calutils"
	"fmt"
	"strconv"
)

func main() {
	var stockcode int = 123
	var enddate string = "2020-12-31"
	var url string = "Code=%d&endDate=%s"
	var target_url string = fmt.Sprint(url, stockcode, enddate) // 根据格式化字符串，将变量stockcode和enddate插入到url中
	fmt.Println(target_url)                                     // 打印出目标url
	// 将字符串转换为int类型
	var s1 string = "k10"
	var n1 int
	var err error
	n1, err = strconv.Atoi(s1) // 该函数返回2个值,将字符串s1转换为int类型，如果转换失败，则返回错误信息
	if err != nil {
		fmt.Println("convert failed:", err)
	} else {
		fmt.Println("convert success:", n1)
	}
	// 将int转换为string类型
	var n2 int = 123
	var s2 string = strconv.Itoa(n2)
	fmt.Printf("int %d convert to string: '%s'", n2, s2)
	// string 转 float
	var s3 string = "123.45"
	var f1 float64
	var err2 error
	f1, err2 = strconv.ParseFloat(s3, 64)
	if err2 != nil {
		fmt.Println("convert failed:", err2)
	} else {
		fmt.Println("convert success:", f1)
	}
	// float 转 string
	var f2 float64 = 123.45
	var s4 string = strconv.FormatFloat(f2, 'f', 2, 64)
	fmt.Printf("float %f convert to string: '%s'\n", f2, s4)

	fmt.Println(calutils.Product(1, 2))

}
