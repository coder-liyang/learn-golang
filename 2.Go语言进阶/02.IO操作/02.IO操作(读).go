package main

import (
	"fmt"
	"io"
	"os"
)

//IO包中最重要的两个接口 Reader和Writer
func main() {
	fileName := "demo.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	} else {
		defer file.Close()
		fmt.Println(file.Name())
	}
	bs := make([]byte, 4, 4)
	//n , err := file.Read(bs)
	//if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Println(n)
	//	fmt.Println(bs)
	//	fmt.Println(string(bs))
	//}
	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("已经读到了文件末尾")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
