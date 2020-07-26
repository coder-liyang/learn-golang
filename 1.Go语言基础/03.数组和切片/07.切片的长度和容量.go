package main

import "fmt"

func main() {
	//var s []int
	//var s = make([]int, 0, 5)
	s := make([]int, 2, 5)
	fmt.Printf("长度:%d, 容量:%d\n", len(s), cap(s))

	//s[0] = 1 //panic: runtime error: index out of range [0] with length 0

	//当容量不够时,会翻倍
	for i:=0; i<10;i++ {
		s = append(s, 1)
		fmt.Printf("长度:%d, 容量:%d\n", len(s), cap(s))
	}
}