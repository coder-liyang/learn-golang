package main

import "fmt"

const c3 = "This is c3"
func main() {
	//常量的类型可选
	const c1 = "nihao"
	fmt.Println(c1)

	const c2 string = "haha"//常量可不使用

	fmt.Println(c3)

	const (
		man = 1
		woman = 2
		unknown = 0
	)
	fmt.Println(man, woman, unknown)

	const (
		cc1 = 100
		cc2 float32 = 100.1
		cc3 float64 = 100.1
		cc4 = 100.1234567
		cc5
		cc6
	)
	fmt.Printf("%T, %d \n", cc1, cc1)
	fmt.Printf("%T, %f \n", cc2, cc2)
	fmt.Printf("%T, %f \n", cc3, cc3)
	fmt.Printf("%T, %f \n", cc4, cc4)
	fmt.Printf("%T, %f \n", cc5, cc5)
	fmt.Printf("%T, %f \n", cc6, cc6)

	const (
		ccc1 = iota
		ccc2 = iota
		ccc3
		ccc4 = iota
		ccc5
		ccc6 = "nihao"
		ccc7
		ccc8 = iota
	)
	fmt.Printf("%T, %d \n", ccc1, ccc1)
	fmt.Printf("%T, %d \n", ccc2, ccc2)
	fmt.Printf("%T, %d \n", ccc3, ccc3)
	fmt.Printf("%T, %d \n", ccc4, ccc4)
	fmt.Printf("%T, %d \n", ccc5, ccc5)
	fmt.Printf("%T, %s \n", ccc6, ccc6)
	fmt.Printf("%T, %s \n", ccc7, ccc7)
	fmt.Printf("%T, %d \n", ccc8, ccc8)

}