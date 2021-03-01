package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func main() {
	//WaitGroup
	//Add()设置要等待的goroutine的数量
	//Done()将数量-1

	wg2.Add(2)
	go printNum()
	go printStr()
	fmt.Println("main() 阻塞")
	wg2.Wait()
	fmt.Println("main()，解除阻塞。。")
}

func printNum() {
	defer wg2.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("printNum: %d\n", i)
	}
}

func printStr() {
	defer wg2.Done()
	for i := 'a'; i < 'w'; i++ {
		fmt.Printf("printNum: %c\n", i)
	}
}
