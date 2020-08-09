package main

import "fmt"

//局部变量: 函数内部定义的变量
//全局变量: 函数外部定义的变量

var g string = "hello"

func main() {
	var p string = "world"

	fmt.Println(g, p)
}