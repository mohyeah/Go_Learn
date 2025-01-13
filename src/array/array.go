package main

import "fmt"

func main() {
	var arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var arr2 = [5]int{}
	// go中, 数组大小是类型的一部分, 故 [10]int 和 [5]int 是不同的类型
	arr3 := [...]float32{1.1, 2.2, 3.3, 4.4, 5.5} // 未指定数组长度, 编译器会自动计算数组长度

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Println(arr2)
	fmt.Println(arr3)
	// 对数组索引位置赋值
	arr2 = [5]int{0: 1, 3: 3}
	fmt.Println(arr2)

	// 多维数组
	var arr23 [2][3]int = [2][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
	}
	fmt.Println(arr23)

	// 访问数组元素
	var salary float32 = arr3[0]
	fmt.Println(salary)
	// 切片
	var slice1 []int = arr1[0:5]
	fmt.Println(slice1)
	// 动态创建切片, 每当加入的元素超过数组长度, 则自动扩容
	var slice2 []int = make([]int, 3, 5) // 创建一个长度为3, 容量为5的切片
	fmt.Printf("slice2: len=%d, cap=%d\n", len(slice2), cap(slice2))
	for i := 1; i < 6; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Printf("slice2: len=%d, cap=%d", len(slice2), cap(slice2))

	var arr4 []int = []int{1, 2, 3, 4, 5}
	// 添加元素
	arr4 = append(arr4, 6, 7)
	// 将arr4的元素复制到arr5
	var arr5 []int
	copy(arr5, arr4)

	// map函数
	var capital_map = map[string]string{"China": "Beijing", "USA": "Washington"}
	fmt.Println(capital_map)
	for country := range capital_map {
		fmt.Println(country, "capital is:", capital_map[country])
	}
	// 删除map中的元素
	delete(capital_map, "USA")
	fmt.Println(capital_map)

	for country, capital := range capital_map {
		fmt.Println(country, "capital is:", capital)
	}

	var arr6 = [5]string{}
	for i := 0; i < len(arr6); i++ {
		fmt.Printf("input No.%d score:", i)
		fmt.Scanf("%s", &arr6[i])
	}
	for key, val := range arr6 {
		fmt.Printf("No.%d score: is %s\n", key, val)
	}

}
