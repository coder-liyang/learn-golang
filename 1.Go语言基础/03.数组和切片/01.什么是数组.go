package main

import "fmt"

func main() {
	//先声明,后赋值
	var value1 [10]int
	value1[0] = 1
	value1[2] = 89
	fmt.Println(value1)

	//声明同时赋值
	var value2 = [10]string{"hello", "world"}
	fmt.Println(value2)

	//byte类型
	var value3 = [3]int{'A', 'B', 'C'} //单引号
	fmt.Println(value3)

	//自动推算数量
	var value4 = [...]string{"fd", "f", "d"}
	fmt.Println(value4)

	//给指定位置赋值
	var value5 = [5]int{2: 10, 3: 333}
	fmt.Println(value5)

	//简短定义
	value6 := [...]int{2, 3, 4, 5, 6, 19: 999}
	fmt.Println(value6)

}
