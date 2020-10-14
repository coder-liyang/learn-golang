package main

import (
	"fmt"
	c "learn-golang/1.Go语言基础/08.结构体/attr"
)

func main() {
	var pc1 c.Pc
	pc1.Color = "black"
	pc1.Name = "联想"
	pc1.Price = 5000
	//pc1.purchasingCost = 4000 //小写的属性无法在包外访问

	fmt.Println(pc1)

	//var pc2 c.pc2 //无法使用
}
