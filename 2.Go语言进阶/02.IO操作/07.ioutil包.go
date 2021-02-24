package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.OpenFile("demo.txt", os.O_RDONLY, os.ModePerm)
	f, _ := ioutil.ReadAll(file)
	fmt.Println("f\t", string(f))

	f2, _ := ioutil.ReadFile("demo.txt")
	fmt.Println("f2\t", string(f2))

	ioutil.WriteFile("demo9.txt", []byte("hello world"), os.ModePerm)

	f3, _ := ioutil.ReadFile("demo9.txt")
	fmt.Println("f3\t", string(f3))

	fileInfo, _ := ioutil.ReadDir("./")
	for k, v := range fileInfo {
		fmt.Println(k, "---", v.Name())
	}
}
