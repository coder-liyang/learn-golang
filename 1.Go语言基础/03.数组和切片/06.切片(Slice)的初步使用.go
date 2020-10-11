package main

import "fmt"

func main() {
	//切片不保存数据,本质上是数的引用

	//定义切片1
	//var s []string
	//定义切片2
	var s []string = make([]string, 0)
	//简写
	//s := make([]string, 0)
	for i := 'a'; i <= 'z'; i++ {
		s = append(s, fmt.Sprintf("%c", i))
	}
	fmt.Println("原始切片:", s)
	//两个变量实际上指向的同一块内存地址
	ss := s

	//按照下标截取,左闭右开
	fmt.Println(s[10:18])
	fmt.Println(s[1:])
	fmt.Println(s[:5])

	s[10] = "hello"
	fmt.Println(s)
	//对原切片的修改,会影响赋值过的其他切片
	fmt.Println(ss)

	var arr [10]int
	for j := 0; j < 10; j++ {
		arr[j] = j + 1
	}
	fmt.Println(arr)

	var s2 = arr[3:5]
	fmt.Println(s2)
	for k, _ := range s2 {
		s2[k] += 100
	}
	fmt.Println(s2)
	fmt.Println(arr)
}
