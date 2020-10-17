package main

import "fmt"

type Bird struct {
	Category string
	Run      func() string
}

func (b *Bird) Fly() {
	fmt.Printf("%s can fly.\n", b.Category)
}
func main() {
	//方法就是包含了接受者的函数,接受者可以是命名类型或者结构体类型的一个值或一个指针
	b1 := Bird{Category: "麻雀"}
	b1.Run = func() string {
		fmt.Printf("%s can run.\n", b1.Category)
		return b1.Category
	}
	b1.Fly()
	fmt.Println(b1.Run())
	b2 := Bird{Category: "鸵鸟"}
	b2.Run = func() string {
		fmt.Printf("%s 跑的很快.\n", b2.Category)
		return b2.Category
	}
	b2.Fly()
	b2.Run()
}
