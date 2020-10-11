package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 100
	ptr = &a
	pptr = &ptr
	fmt.Println(a, ptr, pptr)

	fmt.Println(a, *ptr, **pptr)
}
