package main

import "fmt"

func main() {
	a := 100
	anyParam(a)
	b := "abc"
	anyParam(b)
	c := true
	anyParam2(c)
}

func anyParam(i interface{}) {
	//t1 := i.(string) //不安全,如果类型不是string,则直接panic
	if _, ok := i.(string); ok { //安全,如果类型不是string,则ok的值为false
		fmt.Println("类型为string")
		return
	}
	fmt.Println("未知的类型")
}

func anyParam2(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("是string类型")
	case int:
		fmt.Println("是int类型")
	case bool:
		fmt.Println("是bool类型")
	default:
		fmt.Println("不知道是什么类型")
	}
}
