package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"reflect"
)

func main() {
	//
	// ### 不要忽略错误
	//

	//fileName := "01.什么是错误.go"
	fileName := "no_file.txt"
	f, err := os.Open(fileName)
	fmt.Println(reflect.TypeOf(err))
	//1. 断言底层结构类型并从结构字段获取更多信息
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Path, "文件不存在")
		//return
	} else {
		fmt.Println(f.Name())
	}
	if err != nil {
		fmt.Println("普通错误")
		//return
	}

	addrs, err := net.LookupHost("google.com") //
	//2. 断言底层结构类型，并使用方法获取更多信息
	if err, ok := err.(*net.DNSError); ok {
		fmt.Println(1)
		if err.Timeout() {
			fmt.Println("timeout")
		} else if err.Temporary() {
			fmt.Println("temporary")
		} else {
			fmt.Println(err.Error())
		}
		return
	} else {
		fmt.Println(addrs)
	}
	//3. 直接比较错误
	matches, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
	} else {
		fmt.Println("matches files", matches)
	}
}

type FileError struct {
}

func (c *FileError) Error() string {
	return "文件未找到"
}
