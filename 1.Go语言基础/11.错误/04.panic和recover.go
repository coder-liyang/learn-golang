package main

import "fmt"

//什么时候用错误,什么时候用异常
//异常的场景: 1.空指针的引用 2.下标越界 3.除数为0 4.不应该出现的分支,比如default 5.输入不正确
//其他的基本上都是错误

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("出错了,被recover了")
		}
	}()

	panic("这是一个故意制造的panic")
	//sli := make([]string, 1)
	//sli = append(sli, "aaa")
	//fmt.Println(len(sli), getSlice(sli))
}

func getSlice(s []string) string {
	return s[100]
}
