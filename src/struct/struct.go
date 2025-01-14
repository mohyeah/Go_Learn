package main

import "fmt"

type student struct {
	name    string
	age     int
	gender  string
	subject string
	score   float64
}

// 结构体指针
func printStudent(stu *student) {
	fmt.Printf("student name: %s\n", stu.name)
	fmt.Printf("student age: %d\n", stu.age)
	fmt.Printf("student gender: %s\n", stu.gender)
	fmt.Printf("student subject: %s\n", stu.subject)
	fmt.Printf("student score: %.2f\n", stu.score)
}

// 结构体方法
func (stu *student) selfDes() {
	fmt.Println("I'm", stu.name)
}

func main() {
	// 创建结构体变量
	fmt.Println(student{"张三", 18, "男", "语文", 90.5})
	// 字典赋值
	fmt.Println(student{name: "李四", age: 19, gender: "男", subject: "数学", score: 80.5})
	// 未赋值的字段默认为零值
	fmt.Println(student{name: "王五", age: 20})
	// 访问结构体变量
	var Lily student
	Lily.name = "Lily"
	Lily.age = 18
	fmt.Println(Lily.age)
	printStudent(&Lily)
	Lily.selfDes()

}
