package main

import "fmt"

func main() {
	var v int

	m1 := map[string]int{"a": 11, "b": 22, "c": 33, "d": 44}
	fmt.Println(m1["b"], m1["e"])

	v = m1["b"]
	fmt.Println(v)

	if v, ok := m1["x"]; ok {
		fmt.Println("值存在", v)
	} else {
		fmt.Println("值不存在")
	}
	if v, ok := m1["c"]; ok {
		fmt.Println("值存在", v)
	} else {
		fmt.Println("值不存在")
	}
}
