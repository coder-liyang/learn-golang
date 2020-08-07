package main

import "fmt"

func main() {
	//函数定义时的参数:形参
	//函数调用时的参数:实参

	//函数调用时,函数名必须匹配,实参顺序必须一一对应:顺序/个数/类型

	//可变参数
	sum := myFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 1, 2, 2, 3, 4, 23, 432, 43, 5, 345, 43, 645, 6, 54)
	fmt.Println(sum)

	//值传递和引用传递
	//值传递
	s1 := "hahahahah"
	s2 := myFunc2(s1)
	fmt.Println(s2, s1)
	//引用传递
	s3 := "hehehehe"
	s4 := myFunc3(&s3) //&s3传递地址
	fmt.Println(s4, s3)

}

//可变参数,个数不定,但类型必须都是int
func myFunc(params ...int) int {
	var sum int
	for _, v := range params {
		sum += v
	}
	return sum
}

//值传递
func myFunc2(s string) string {
	s += "函数内修改了值"
	return s
}

//接收的是指针
func myFunc3(s *string) string {
	*s += "函数内修改了值"
	return *s
}
