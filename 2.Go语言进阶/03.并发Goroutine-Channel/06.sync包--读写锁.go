package main

import (
	"fmt"
	"sync"
	"time"
)

var wg3 *sync.WaitGroup
var rwMutex *sync.RWMutex

/*
读锁不能阻塞读锁
读锁需要阻塞写锁，直到所有读锁都释放
写锁需要阻塞读锁，直到所有写锁都释放
写锁需要阻塞写锁
*/
func main() {
	wg3 = new(sync.WaitGroup)
	rwMutex = new(sync.RWMutex)
	//读锁
	wg3.Add(2)
	go read("A")
	go read("B")
	wg3.Wait()
	fmt.Println("=====================")

	//写锁
	wg3.Add(4)
	go write("AA")
	go read("BB")
	go write("CC")
	go read("DD")
	wg3.Wait()
}

func read(name string) {
	defer wg3.Done()
	fmt.Printf("%s开始读\n", name)
	//读锁
	rwMutex.RLock()
	fmt.Printf("%s正在读\n", name)
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
	fmt.Printf("%s读结束\n", name)
}

func write(name string) {
	defer wg3.Done()
	fmt.Printf("%s开始写\n", name)
	rwMutex.Lock()
	fmt.Printf("%s正在写\n", name)
	time.Sleep(1 * time.Second)
	rwMutex.Unlock()
	fmt.Printf("%s写结束\n", name)
}
