package main

import "fmt"

func main() {
	a := [...]string{"b", "g", "a", "d", "c", "f", "e"}
	fmt.Println(a)

	//冒泡
	//for k, v := range a {
	//	if
	//}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)

}
