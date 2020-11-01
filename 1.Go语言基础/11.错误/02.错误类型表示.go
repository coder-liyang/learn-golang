package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	//fileName := "01.什么是错误.go"
	fileName := "no_file.txt"
	f, err := os.Open(fileName)
	fmt.Println(reflect.TypeOf(err))
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Path, "文件不存在")
		return
	}
	if err != nil {
		fmt.Println("普通错误")
		return
	}
	fmt.Println(f.Name())
}

type FileError struct {
}

func (c *FileError) Error() string {
	return "文件未找到"
}
