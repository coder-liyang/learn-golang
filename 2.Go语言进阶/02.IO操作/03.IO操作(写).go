package main

import "os"

func main() {
	fileName := "./demo3.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	//将字节数组写入文件
	_, err = file.Write([]byte("file.Write:hello world\n"))
	if err != nil {
		panic(err.Error())
	}
	//将字符串写入文件
	_, err = file.WriteString("file.WriteString:hello world\n")
	if err != nil {
		panic(err.Error())
	}
}
