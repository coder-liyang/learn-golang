package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go r(ch)
	//需要判断通道是否关闭
	for {
		data, ok := <-ch
		if ok {
			fmt.Println("从通道中取出的值是:", data)
		} else {
			fmt.Println("通道关闭,从通道中取出的值是:", data)
			break
		}
	}

	ch1 := make(chan int)
	go r(ch1)
	//无需判断通道是否关闭
	for v := range ch1 {
		fmt.Println("从通道中取出的值是:", v)
	}
}

func r(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i
	}
	close(ch)
}
