package main

import (
	"fmt"
	"os"
)

//错误指的是可能出现问题的地方出现了问题
//而异常指的是不应该出现问题的地方出现了问题

func main() {
	f, err := os.Open("no_file.txt")
	if err != nil {
		fmt.Println("文件打开错误")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
