package main

import (
	"fmt"
	"os"
)

func main() {
	if ReadFile() {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAIL")
	}
}

func ReadFile() bool {
	file, err := os.Open("05.txt")
	if err != nil {
		fmt.Println(err)
		defer file.Close()
		return false
	}

	s := []byte{}
	n, err:=file.Read(s)
	fmt.Println(n)
	return true
}