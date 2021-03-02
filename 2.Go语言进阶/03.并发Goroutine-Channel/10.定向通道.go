package main

import "fmt"

func main() {
	ch1 := make(chan int) //双向 可读可写
	//ch2 := make(<-chan int) //单向 只读 ch2 <-100 报错
	//ch3 := make(chan<- int) //单向 只写 <-ch3 报错
	//单向通道一般用在函数的形参上
	go c1(ch1)
	fmt.Println(c2(ch1))
	close(ch1)
}

//只写
func c1(ch chan<- int) {
	ch <- 100
}

//只读
func c2(ch <-chan int) int {
	return <-ch
}
