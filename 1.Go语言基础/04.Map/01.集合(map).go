package main

import "fmt"

func main() {
	//map是无序的
	//map的长度是不固定的,和slice一样,也是引用类型
	//map可以使用len函数
	//map的key是所有可以比较的类型

	m1 := make(map[string]string)
	fmt.Println(m1)

	m1["a"] = "aaa"
	m1["b"] = "bbb"
	m1["c"] = "ccc"
	m1["d"] = "ddd"
	fmt.Println(m1, len(m1))

	for k,v:=range m1{
		fmt.Printf("%s->%s\n", k, v)
	}

	if _,ok := m1["b"]; ok == true {
		fmt.Println("有值")
	} else {
		fmt.Println("没有值")
	}
}
