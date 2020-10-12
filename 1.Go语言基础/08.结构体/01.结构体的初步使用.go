package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{
		name: "小明",
	}
	p1.age = 18
	fmt.Println(p1)

	p2 := Person{"小黑", 5}
	fmt.Println(p2.name, p2.age)

	p3 := new(Person)
	p3.name = "小红"
	p3.age = 16
	fmt.Println(p3, *p3)

}
