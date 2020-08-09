package main

import "fmt"

func main() {
	//一个函数可以没有返回值,可以有一个返回值,可以没有返回值
	fmt.Println(fun1("aaa", "bbb"))
	//可以忽略返回值
	a, _ := fun1("qqqq", "eeeeee")
	fmt.Println(a)
}

func fun1(a, b string) (string, string) {
	b, a = a, b
	return a, b
}
