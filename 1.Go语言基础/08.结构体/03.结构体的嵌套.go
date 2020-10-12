package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type Cat struct {
	Animal
	run bool
	int
}

type Bird struct {
	Animal
	fly bool
}

func main() {
	c1 := Cat{
		Animal: Animal{
			name: "小喵",
			age:  5,
		},
		run: true,
		int: 8, //所有的内置类型和自定义类型都是可以作为你命字段的
	}
	//c1.name叫做提升字段
	fmt.Println(c1.name, c1.Animal.name, c1.int)
}
