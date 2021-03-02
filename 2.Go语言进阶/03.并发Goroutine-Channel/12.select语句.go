package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//select 是 Go 中的一个控制结构。select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。
	//如果没有case可运行，它将阻塞，直到有case可运行。

	ch1 := make(chan string)
	ch2 := make(chan string)

	rand.Seed(time.Now().UnixNano())
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		ch1 <- "ch1"
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		ch2 <- "ch2"
	}()

	//fmt.Println(rand.Intn(5))
	select {
	case data1 := <-ch1:
		fmt.Println("ch1中取出的值是:", data1)
	case data2 := <-ch2:
		fmt.Println("ch2中取出的值是:", data2)
	case <-time.After(3 * time.Second):
		fmt.Println("time.After执行了")
		//default: //如果ch1和ch2都没有值,默认走这里
		//	fmt.Println("ch1和ch2都没有值")
	}
}
