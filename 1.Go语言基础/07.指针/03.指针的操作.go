package main

import "fmt"

func main() {
	//获取一个指针意味着访问指针所指向变量的值,语法是 *a
	var a string = "aaa"
	var b *string
	b = &a

	fmt.Println(a, b)
	fmt.Println(&a, *b)
	//操作指针,改变a的值
	*b = "bbb"
	fmt.Println(a)
	//使用指针传递函数的参数
	var a1 = 10
	fmt.Println(a1)
	change(&a1)
	fmt.Println(a1)
	//将数组的指针传递给函数(这不是一个常规的做法)
	a2 := [3]int{1, 2, 3}
	fmt.Println(a2)
	changeArr(&a2)
	fmt.Print(a2)
	//用切片实现上面的效果
	s2 = []{}
}

func change(a *int) {
	*a++
}

func changeArr(a *[3]int) {
	(*a)[0]++
	(*a)[1]++
	(*a)[2]++
}

func changeSlice(s []int) {
	for i := 0; i< len(s); i++  {
		s[i]++
	}
}
