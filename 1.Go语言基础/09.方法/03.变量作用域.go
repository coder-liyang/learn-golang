package main

import "fmt"

//变量可以在三个地方声明
//1.函数内定义的变量称为局部变量
//2.函数外定义的变量称为全局变量
//3.函数定义中的变量称为形式参数

//声明全局变量
var g int

func main() {
	//声明局部变量
	var a, b int

	a = 10
	b = 20
	g = a + b
	fmt.Println(g)
	fmt.Println(a)
	fmt.Println(pa(a))
	//函数内的a只是一份拷贝,修改了函数内a的值,不会影响原始的变量
	fmt.Println(a)
	//但是如果传递的是指针,就会影响原始变量了
	fmt.Println(pa2(&a))
	fmt.Println(a)
}

func pa(a int) int {
	a++
	return a
}

func pa2(a *int) int {
	*a++
	return *a
}
