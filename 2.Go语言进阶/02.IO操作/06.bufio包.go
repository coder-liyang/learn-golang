package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//读数据相关
	//read()
	//写数据相关
	write()
}

func read() {
	file, _ := os.Open("demo.txt")
	reader := bufio.NewReader(file)
	p := make([]byte, 2)
	reader.Read(p)
	fmt.Println(string(p))

	p2, _, _ := reader.ReadLine()
	fmt.Println(string(p2))

	for {
		s1, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(s1)
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(s1)
	}

}

func write() {
	file, _ := os.OpenFile("demo8.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)
	len, err := writer.WriteString("abcdefg")
	//len, err := writer.Write([]byte{'A', 'B'})
	if err != nil {
		panic(err.Error())
	}
	writer.Flush()
	fmt.Printf("写入成功,写入长度:%d\n", len)

}
