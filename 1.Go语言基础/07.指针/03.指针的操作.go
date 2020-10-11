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
	s2 := make([]int, 0)
	s2 = append(s2, 0, 1, 2, 5, 19, 31)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(s2)
	changeSlice(s2)
	fmt.Println(s2)
	fmt.Println("!!!!!!!!!!!!!!!!!")
	change2()
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
	for i := 0; i < len(s); i++ {
		s[i]++
	}
}

func change2() {
	slice1 := make([]string, 0)
	slice1 = append(slice1, "a", "b", "c", "d", "e", "f")
	fmt.Println(slice1)
	array := [6]*string{}
	for i := 0; i < len(slice1); i++ {
		array[i] = &slice1[i]
	}
	fmt.Println(array)
	slice1[2] = "haha" //如果切片改了,那么数组也跟着改
	for j := 0; j < len(slice1); j++ {
		fmt.Print(array[j], "=")
		fmt.Print(*array[j], " ")
	}
}
