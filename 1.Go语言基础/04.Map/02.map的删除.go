package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"] = "aaa"
	fmt.Println(m)

	m2 := map[string]string{"b":"bb", "c":"cc"}
	fmt.Println(m2)
	delete(m2, "b")
	fmt.Print(m2)
}