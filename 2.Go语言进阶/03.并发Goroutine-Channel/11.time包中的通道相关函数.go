package main

import (
	"fmt"
	"time"
)

func main() {
	//新建一个定时器 time.NewTimer()
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("类型:%T\n", timer)
	fmt.Println(time.Now())
	ch2 := timer.C //此代码会阻塞3秒
	fmt.Println(<-ch2)

	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2结束")
	}()
	time.Sleep(3 * time.Second)
	//计时器停止 timer.Stop
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2停止")
	}

	fmt.Println("=======time.After()=======")
	//time.After()表示多长时间长的时候后返回一条time.Time类型的通道消息。但是在取出channel内容之前不阻塞，后续程序可以继续执行。
	ch3 := time.After(3 * time.Second)
	fmt.Println(time.Now())
	time2 := <-ch3
	fmt.Println(time2)
}
