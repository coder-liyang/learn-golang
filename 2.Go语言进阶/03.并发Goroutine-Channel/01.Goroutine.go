package main

import (
	"fmt"
	"time"
)

func main() {

	//go hello()
	////如果不Sleep,有可能先输出main,再输出hello,有可能先输出hello,在输出main,有可能只输出main
	//time.Sleep(1 * time.Second)
	//fmt.Println("This is main()")

	//
	go numbers()
	go alphabets()
	time.Sleep(3 * time.Second)
	fmt.Println("This is main()")
}

func hello() {
	fmt.Println("This is hello()")
}

func numbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Microsecond)
		fmt.Print(i)
	}
}

func alphabets() {
	for i := 'a'; i < 'g'; i++ {
		time.Sleep(300 * time.Microsecond)
		fmt.Printf("%c", i)
	}
}
