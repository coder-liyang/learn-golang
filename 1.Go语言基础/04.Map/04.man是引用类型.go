package main

import "fmt"

func main() {
	m1 := map[string]string{"a":"aa","b":"bb","c":"cc"}
	fmt.Println(m1)
	m2 := m1
	fmt.Println(m2)

	println()

	delete(m1, "b")
	m1["d"] = "dd"
	fmt.Println(m1, m2)
}