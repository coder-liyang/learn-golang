package main

import (
	"fmt"
)

//临界资源: 指并发环境中多个进程/线程/协程共享的资源。

var num = 5

//go run --race 03.临界资源安全问题.go
func main() {
	//一个简单的示例
	//a := 1
	//go func() {
	//	a = 2
	//	fmt.Println("子goroutine:", a)
	//}()
	//a = 3
	//time.Sleep(1 * time.Second)
	//fmt.Println("主goroutine:", a)

	//卖票的示例
	//ticket()
	//time.Sleep(1 * time.Second)

	//解决临界安全的问题
	ticket2()
}

func ticket() {
	go sale("A")
	go sale("B")
	go sale("C")
	go sale("D")
}

func sale(name string) {
	for {
		if num <= 0 {
			break
		}
		fmt.Printf("%s查询到%d张,可以出票\n", name, num)
		num--
		fmt.Printf("%s卖出1张,还剩%d张\n", name, num)
	}
	fmt.Printf("%s窗口售票结束\n", name)
}

func ticket2() {

}

func sale2(name string) {

}
