package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var num1 int64 = 10
var mutex1 sync.Mutex
var wg1 sync.WaitGroup

func main() {
	wg1.Add(4)
	go sale1("A")
	go sale1("B")
	go sale1("C")
	go sale1("D")
	wg1.Wait()
}

func sale1(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg1.Done()
	for {
		mutex1.Lock()
		if num1 <= 0 {
			mutex1.Unlock()
			fmt.Printf("%s:当前余票%d,售票结束\n", name, num1)
			return
		}
		fmt.Printf("%s:当前余票%d,可以继续售票\n", name, num1)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		num1--
		mutex1.Unlock()
		fmt.Printf("\t%s:卖出一张,当前余票%d\n", name, num1)
	}
}
