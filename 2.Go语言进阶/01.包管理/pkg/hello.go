package pkg

import "fmt"

func Hello() {
	fmt.Println("此函数可以被导出")
}

func hello() {
	fmt.Println("此函数为私有,不可导出")
}
