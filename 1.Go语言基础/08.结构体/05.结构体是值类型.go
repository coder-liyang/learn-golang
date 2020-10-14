package main

import "fmt"

func main() {
	type name struct {
		firstName string
		lastName  string
	}

	name1 := name{
		firstName: "张",
		lastName:  "三",
	}

	name2 := name{
		firstName: "李",
		lastName:  "四",
	}

	name3 := name{
		firstName: "张",
		lastName:  "三",
	}

	//结构体是值类型,如果每个字段具有可比性,则是可比较的.如果他们对应的字段相等,则认为两个结构体变量是相等的
	fmt.Println(name1 == name2, name1 == name3)

	type data struct {
		list []string
	}

	data1 := data{list: []string{"a", "b"}}
	data2 := data{list: []string{"a", "b"}}
	fmt.Println(data1, data2)
	//如果结构体中包含的字段是不可比较的,那么结构体变量是不可以比较的
	//fmt.Println(data1 == data2) //invalid operation: data1 == data2 (struct containing []string cannot be compared)

	type Person struct {
		name name
		age  int
	}

	p1 := Person{
		name: name{"张", "三"},
		age:  10,
	}
	p2 := Person{
		name: name{"张", "三"},
		age:  10,
	}
	fmt.Println(p2, p2)
	//如果嵌套的结构体是可以比较的,那么这个结构体仍然是可以比较的
	fmt.Println(p1 == p2)
}
