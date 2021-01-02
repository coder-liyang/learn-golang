package main

import "learn-golang/2.Go语言进阶/01.包管理/pkg"

import (
	. "fmt"
	time2 "time" //取了一个别名
)

//1.main包:go语言的入口main()函数所在的包叫做main,main包使用其他包的代码,需要导入(import)
//2.package: 同一个包下的所有.go文件的第一行添加包定义,一标记该文件所属的包
//	一个目录下的同级文件属于同一个包
//	报名建议设置为目录名
//	main包为程序入口包,其他包不可用
//	包可以嵌套
//	首字母大写的函数可以被导出
//3.导入使用import关键字
//	可以单个导入也可以批量导入
//	'.'操作符 调用时可以省略包前缀
//	'_'操作符 只执行包内的init方法
//	别名 通过别名使用包内的函数
func main() {
	pkg.Hello()

	Println("由于导入的时候使用了'.'操作符,因此这里可以直接使用Println")

	Println(time2.Now().Format("2006-01-02 15:04:05"))
}
