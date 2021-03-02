package main

import (
	"fmt"
	"time"
)

/*
1.channel用于goroutine，传递消息的。
2.通道，每个都有相关联的数据类型,nil chan，不能使用，类似于nil map，不能直接存储键值对
3.使用通道传递数据：<-
	chan <- data,发送数据到通道。向通道中写数据
	data <- chan,从通道中获取数据。从通道中读数据
4.阻塞：
	发送数据：chan <- data,阻塞的，直到另一条goroutine，读取数据来解除阻塞
	读取数据：data <- chan,也是阻塞的。直到另一条goroutine，写出数据解除阻塞。
5.本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作。
6.通道是goroutine之间的连接，所以通道的发送和接收必须处在不同的goroutine中。
*/
func main() {
	var a chan int
	if a == nil {
		fmt.Println("通道a是nil,,无法使用,需要先创建通道...")
	}
	a = make(chan int)
	//简短声明
	//b := make(chan int)
	fmt.Printf("通道a的类型是%T,地址是%p\n", a, a)
	//channel是引用类型的数据，在作为参数传递的时候，传递的是内存地址。
	go test1(a)
	//这里会阻塞,一直等到有test1()把数据放入通道
	data := <-a
	fmt.Println(data)

	//b := make(chan int)
	//<-b 只出不如会死锁
	//b <- 1 只出不入会死锁
}

func test1(ch chan int) {
	time.Sleep(1 * time.Second)
	ch <- 100
	fmt.Printf("函数内通道的类型是%T,地址是%p\n", ch, ch)
}
