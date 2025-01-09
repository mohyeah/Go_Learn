package main

import "fmt"

func main() {
	var stockcode int = 123
	var enddate string = "2020-12-31"
	var url string = "Code=%d&endDate=%s"
	var target_url string = fmt.Sprint(url, stockcode, enddate) // 根据格式化字符串，将变量stockcode和enddate插入到url中
	fmt.Println(target_url)                                     // 打印出目标url
}
