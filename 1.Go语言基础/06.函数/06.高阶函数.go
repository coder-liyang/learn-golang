package main

import "fmt"

func main() {
	//函数也是一种数据类型
	var f1 = fun10
	var s = "fdfs"
	fmt.Println(s)
	var ss string = "kjkjkjj"
	fmt.Println(ss)
	fmt.Printf("%s,%T\n", f1, f1)

	f2 := fun10(123, 123423)
	fmt.Println(f2)
	//匿名函数
	f3 := func() string {
		return "666"
	}
	fmt.Println(f3())

	//回调函数,将一个函数fun2作为函数fun1的参数,那么fun2叫回调函数,fun1叫高阶函数
	add := func11(100, 200, add)
	fmt.Println(add)

	fmt.Println(increment()())
	fmt.Println(increment()())

	res1 := increment()
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

}

func fun10(num1 int, num2 int) int {
	return num1 + num2
}

//需要传入一个匿名函数,由匿名函数决定如何计算
func func11(num1 int, num2 int, callback func(num1 int, num2 int) int) int {
	return callback(num1, num2)
}
func add(num1, num2 int) int {
	return num1 + num2
}

func increment() func() int{
	//1.定义一个局部变量
	i := 0
	//2.定义一个匿名函数,给变量加1,并返回
	f := func() int {
		i++
		return i
	}
	return f
}