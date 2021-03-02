package main

import (
	"fmt"
	"time"
)

//缓冲通道就是指一个通道，带有一个缓冲区。发送到一个缓冲通道只有在缓冲区满时才被阻塞。类似地，从缓冲通道接收的信息只有在缓冲区为空时才会被阻塞。

func main() {
	ch1 := make(chan int, 4)
	fmt.Println("通道长度:", len(ch1), "通道容量:", cap(ch1))

	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
		fmt.Println("子协程运行结束,通道关闭")
	}(ch1)

	for data := range ch1 {
		time.Sleep(1 * time.Second)
		fmt.Println("取出来的内容是:", data)
	}
}
