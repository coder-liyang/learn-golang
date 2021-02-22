package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, _ := os.OpenFile("demo.txt", os.O_RDWR, 0)
	defer file.Close()

	//bs := []byte{0}
	//bs := make([]byte, 1)
	bs := make([]byte, 2, 2)
	file.Read(bs)
	fmt.Println(string(bs))

	/*
		Seek(offset, whence)
		offset:偏移量
		whence:相对位置
		file.Seek(2, io.SeekStart) 相对于开始位置,偏移两个字节
	*/
	file.Seek(2, io.SeekStart) //io.SeekStart == 0
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(2, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(-1, io.SeekEnd)
	file.Read(bs)
	fmt.Println(string(bs))
	//    file.WriteString("ABC") 写入到尾部

	fmt.Println("=========================")

	//断点续传,将demo.txt复制到demo7.txt

	srcFile := "./demo.txt"
	destFile := "./demo7.txt"
	tempFile := destFile + ".temp.txt"

	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer file1.Close()
	defer file2.Close()

	// 1.读取临时文件中的数据
	file3.Seek(0, io.SeekStart)
	bs2 := make([]byte, 2, 2)
	n1, err := file3.Read(bs2)
	//if err != nil {
	//	//panic(err.Error())
	//}
	fmt.Println(n1)
	countStr := string(bs2[:n1])
	fmt.Println(countStr)

	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	//2.设置读写的偏移量
	file1.Seek(count, 0)
	file2.Seek(count, 0)
	data := make([]byte, 2, 2)
	n2 := -1            //读取的数据量
	n3 := -1            //写出的数据量
	total := int(count) //读取的数据总量

	for {
		//3.读取数据
		n2, err = file1.Read(data)
		if err == io.EOF {
			fmt.Println("文件复制完毕")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		//将数据写入到文件
		n3, _ = file2.Write(data[:n2])
		total += n3
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		if total > 20 {
			panic("模拟复制失败")
		}
	}
}
