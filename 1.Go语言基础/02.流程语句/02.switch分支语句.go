package main

import "fmt"

func main() {
	var a interface{}

	a = 100
	a = []string{"fd","fdf"}

	switch x := a.(type) {
	case nil:
		fmt.Printf("类型是:%T", x)
	case int:
		fmt.Printf("类型是:%T", x)
	case bool:
		fmt.Printf("类型是:%T", x)
	case float64:
		fmt.Printf("类型是:%T", x)
	case string:
		fmt.Printf("类型是:%T", x)
	case []string:
		fmt.Printf("类型是:%T", x)
	default:
		fmt.Println("未知类型")
	}
}
