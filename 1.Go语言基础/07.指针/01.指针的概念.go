package main

import "fmt"

func main() {
	var b = 10
	//go语言的取地址符是&,放到一个变量前,就会返回变量的内存地址
	fmt.Printf("变量b的地址是:%x\n", &b)
}
