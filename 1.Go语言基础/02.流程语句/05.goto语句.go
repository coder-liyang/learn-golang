package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
		if i == 6 {
			goto theEnd
		}
	}

	theEnd:
		fmt.Println("执行结束")
}
