package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//临界资源: 指并发环境中多个进程/线程/协程共享的资源。

var num = 5

var num2 = 5
var wg sync.WaitGroup
var mutex sync.Mutex

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
	//存在临界资源安全问题
	//ticket()
	//time.Sleep(1 * time.Second)

	//解决临界安全的问题
	ticket2()
	/*
		在Go的并发编程中有一句很经典的话：不要以共享内存的方式去通信，而要以通信的方式去共享内存。
		在Go语言中并不鼓励用锁保护共享状态的方式在不同的Goroutine中分享信息(以共享内存的方式去通信)。
		而是鼓励通过channel将共享状态或共享状态的变化在各个Goroutine之间传递（以通信的方式去共享内存），
		这样同样能像用锁一样保证在同一的时间只有一个Goroutine访问共享状态。
	*/
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

//解决了临界资源安全问题
func ticket2() {
	wg.Add(4)
	go sale2("AA")
	go sale2("BB")
	go sale2("CC")
	go sale2("DD")
	wg.Wait()
}

func sale2(name string) {
	//rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if num2 <= 0 {
			mutex.Unlock()
			break
		}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("%s查询到%d张,可以出票\n", name, num2)
		num2--
		fmt.Printf("%s卖出1张,还剩%d张\n", name, num2)
		mutex.Unlock()
	}
	fmt.Printf("%s窗口售票结束\n", name)
}
