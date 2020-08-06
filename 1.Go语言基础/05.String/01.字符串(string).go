package main

import (
	"fmt"
	"strings"
)

func main() {
	//go中的字符串是一个字节的切片,可以通过将其内容封装在""中来创建字符串,Go中的字符串是unicode兼容的,并且是utf-8编码的
	str := "hello world"
	fmt.Println(str)

	for k, v := range str{
		fmt.Println(k, v, fmt.Sprintf("%c", v))
	}
	//String包
	a := "h"
	if strings.Contains(str, a) {
		fmt.Printf("字符串%s中包含%s", str, a)
	} else {
		fmt.Printf("字符串%s中不包含%s", str, a)
	}
}
