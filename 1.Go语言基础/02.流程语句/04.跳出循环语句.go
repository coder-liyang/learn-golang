package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		if i == 5 {
			continue
		}
		if i == 7 {
			break
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
