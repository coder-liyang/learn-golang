package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//copyFile1("./demo.txt", "./demo4.txt")
	//copyFile2("./demo.txt", "./demo5.txt")
	//io包中除了Copy,还提供了CopyN和CopyBuffer方法
	//无论哪个Copy方法,最终都是有copyBuffer实现的
	copyFile3("./demo.txt", "./demo6.txt")
}

//通过Read()和Write()方法,边读边写,实现文件的复制,块的大小会影响程序的性能
func copyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("拷贝出错")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}

//通过io包下的copy实现
func copyFile2(srcFile, dstFile string) (int64, error) {
	f1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	f2, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f1.Close()
	defer f2.Close()
	return io.Copy(f2, f1)
}

//通过ioutil包
func copyFile3(srcFile, dstFile string) (int, error) {
	f1, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(dstFile, f1, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return len(f1), nil
}
