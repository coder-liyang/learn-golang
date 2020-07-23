package main

import "fmt"

func main() {
	value := [...]int{0: 0, 1: 11, 20: 22}
	fmt.Println(value)

	for i := 0; i < len(value); i++ {
		fmt.Printf("第%d个元素:%d\n", i, value[i])
	}

	for j, v := range value {
		fmt.Printf("for-range-第%d个元素:%d\n", j, v)
	}
}
