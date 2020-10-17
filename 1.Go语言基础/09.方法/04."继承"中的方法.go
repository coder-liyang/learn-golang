package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi,my name is %s.\n", h.name)
}

type Student struct {
	Human
	School string
}

type Employee struct {
	Human
	company int
}

//重写"父类"方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi,my name is %s. I worked hard.\n", e.name)
}

func main() {
	s := Student{
		Human: Human{
			name: "张三",
			age:  20,
		},
		School: "清华大学",
	}
	s.SayHi()

	e := Employee{
		Human: Human{
			name: "李四",
			age:  30,
		},
		company: 10000,
	}
	e.SayHi()
}
