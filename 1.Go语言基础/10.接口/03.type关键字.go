package main

import "fmt"

//type的用法
//1.定义结构体
type Room struct {
}

//2.定义接口
type Home interface {
}

//3.定义其他新类型 type 新类型名 类型
type newString string
type newInt int

//4.定义函数类型
type customFun func(int, int) int

//5. 类型别名, 非本地类型不能定义方法
type age = int

func main() {
	var a int = 10
	var b newInt = 20
	//c := a+b //报错,由于类型不同,不能进行加法运算,虽然本质上都是int
	//a = b //赋值一样报错
	c := a + int(b)
	fmt.Println(c)
	f := plus()
	//调用回调函数
	fmt.Println(f(5, 6))
	//使用别名
	var aa int = 666
	var bb age = 37
	fmt.Printf("bb的类型是:%T\n", bb)
	fmt.Println(aa + bb)
}

func plus() customFun {
	c := func(a, b int) int {
		return a + b
	}
	return c
}
