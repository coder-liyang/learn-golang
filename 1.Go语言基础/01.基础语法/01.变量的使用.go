package main

import (
	"fmt"
)

var globalString = "This is globalString"

func main() {
	fmt.Println(globalString)
	//#声明变量
	//指定变量类型,声明时如果不带值,则默认为当前类型的零值
	var name string
	fmt.Printf("%T, %s \r\n", name, name)
	name = "abcd"
	fmt.Printf("%T, %s \r\n", name, name)

	var n string = "iiiii"
	fmt.Println(n)

	//根据值自行判断类型
	var name1 = 3131321
	fmt.Printf("%T, %d \r\n", name1, name1)

	//根据值自行判断类型,此变量不可以被声明过,多个变量同时声明时,至少保证有一个新变量,只可用于函数体内
	name2 := "kjkjk"
	fmt.Printf("%T, %s \r\n", name2, name2)

	//#多变量声明
	var a, b, c	= 1, 2, "b"
	fmt.Println(a,b,c)

	var a1, b1, c1 string
	a1 = "a"
	b1 = "吧"
	c1 = "C"
	fmt.Println(a1, b1, c1)

	var (
		a2 string = "nihao"
		b2 int64
		c2 float32
	)
	fmt.Println(a2, b2, c2)
}
