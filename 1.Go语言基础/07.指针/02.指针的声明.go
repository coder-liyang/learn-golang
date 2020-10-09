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
	fmt.Println("====================")
	func2()
	fmt.Println("+++++++++++++++++++")
	func3()
}


func func2() {
	var aaa int = 101
	var bbb *int
	bbb = &aaa

	fmt.Printf("aaa的值是:%x,地址是:%x\n", aaa, &aaa)
	fmt.Printf("bbb的值是:%x\n", bbb)
	fmt.Printf("bbb所指向的值是:%d\n", *bbb)
}

func func3() {
	type name int8
	type first struct {
		a int
		b bool
		name
	}
	fmt.Println(first{
		100, true, 10,
	})
	a := new(first)
	a.a = 200
	a.b = false
	a.name = 30
	fmt.Println(a, *a, "\n", a.a, (*a).a, &a)
	fmt.Println("---------")
	var b *first
	b = a
	fmt.Println(b.a, &b, (*b).name)

	//空指针的值为nil
	var ptr *int
	var ptr2 *string
	fmt.Println(ptr, ptr2)
}