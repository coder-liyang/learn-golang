package main

import "fmt"

//声明指针,*T是指针变量的类型,它指向T类型的值
func main() {

	var p1 *int
	var p2 *float64
	var p3 *string

	var v1 = 100
	var v2 = 100.1
	var v3 = "abccde"

	p1 = &v1
	p2 = &v2
	p3 = &v3

	fmt.Println(p1, p2, p3)
	fmt.Printf("指针的值分别为:%d, %.2f, %s\n", *p1, *p2, *p3)
}
