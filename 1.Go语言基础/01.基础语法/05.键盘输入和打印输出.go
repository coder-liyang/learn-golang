package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int
	var y int
	fmt.Println("请输入一个整数")
	fmt.Scanln(&x)
	fmt.Println("请再输入一个整数")
	fmt.Scanln(&y)

	fmt.Printf("%d+%d=%d\r\n", x, y, x+y)

	fmt.Println("请输入一个字符串")
	reader := bufio.NewReader(os.Stdin)
	str1, _ := reader.ReadString('\n')
	fmt.Printf("读到的数据是:%s", str1)
}