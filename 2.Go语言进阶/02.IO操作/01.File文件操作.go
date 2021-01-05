package main

import (
	"fmt"
	"os"
)

//FileInfo中定义了File信息相关方法
//操作权限,不需要创建时,可以设置为0
func main() {
	//文件信息
	fileInfo, err := os.Stat("./demo.txt")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size())
	fmt.Println("最后修改时间:", fileInfo.ModTime())
	fmt.Println("权限:", fileInfo.Mode())
	//文件操作
	file, err := os.Create("demo2.txt")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(file.Name(), "创建成功")
	}
	file.Close()
	//虽然这里是追加,但是由于上面的Create每次都会清空文件,所以看不出来效果
	file2, err := os.OpenFile("demo2.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err.Error())
	} else {
		_, err = file.WriteString("hello string\n")
		if err != nil {
			panic(err.Error())
		}
		_, err = file.Write([]byte("hello byte\n"))
		if err != nil {
			panic(err.Error())
		}

		file2.Close()
	}
	//删除文件
	//_ = os.Remove("demo2.txt")
}
