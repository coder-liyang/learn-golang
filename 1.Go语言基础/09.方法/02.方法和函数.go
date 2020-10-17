package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func SaySalary(e Employee) {
	fmt.Println(e.Salary)
}

func (e *Employee) SaySalary() {
	fmt.Println(e.Salary)
}
func main() {
	e1 := Employee{
		Name:   "张三",
		Salary: 10000,
	}
	//使用函数实现
	SaySalary(e1)
	//使用方法实现
	e1.SaySalary()
	//函数不可以重名,在不同的类型中的方法可以重名
	//用方法更像是面向对象变成
}
