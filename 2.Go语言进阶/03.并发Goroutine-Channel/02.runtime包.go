package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())

	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println("goroutine")
	//	}
	//}()
	//for i := 0; i < 5; i++ {
	//	//让出时间片,让别的协程先执行
	//	runtime.Gosched()
	//	fmt.Println("main...")
	//}

	go func() {
		fmt.Println("goroutine开始")
		fun()
		fmt.Println("goroutine结束")
	}()
	//主协程阻塞3秒钟
	time.Sleep(3 * time.Second)
	fmt.Println("main()结束")
}

func fun() {
	defer fmt.Println("defer...")
	//return: 终止此函数
	//runtime.Goexit(): 终止所在的协程
	runtime.Goexit()
	fmt.Println("fun 函数")
}
