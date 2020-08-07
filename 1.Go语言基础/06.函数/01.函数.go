package main

import "fmt"

func main() {
	fmt.Println(funcName(2, 3))
}
// func 方法名(参数1 类型, 参数2 类型) 返回值类型 {
//		方法体
//		return 指定类型
// }
func funcName(num1 int, num2 int) int {
	return num1 + num2
}
