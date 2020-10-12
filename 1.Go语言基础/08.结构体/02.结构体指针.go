package main

import "fmt"

type Book struct {
	name     string
	size     int
	category string
}

func main() {
	var sanguo2 Book
	sanguo2.name = "三国"
	sanguo2.size = 1000
	sanguo2.category = "小说"
	fmt.Println(sanguo2)

	//结构体指针,可以储存结构体的地址
	var sanguo *Book
	sanguo = &sanguo2
	fmt.Println(sanguo)

}

func (s Book) String() string {
	return s.name
}
