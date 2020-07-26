package main

import "fmt"

//go中的数组是值类型,不是引用类型
func main() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("a:", a)

	b := a
	fmt.Println("b:", b)

	b[0] = "aaa"
	//改变b中元素的值,不会影响a
	fmt.Println(a, b)
}
