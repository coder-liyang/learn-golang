package main

import "fmt"

func main() {
	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s1)

	s2 := make([]int, 0, 5)
	s2 = append(s2, 100, 200, 300)
	//把长的拷贝到短的,超过len的部分会忽略
	copy(s2, s1)
	fmt.Println(s2)

	var s3 []int
	s3 = append(s3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	var s4 = make([]int, 5, 5)
	s4[2] = 200
	s4[4] = 400
	//把短的拷贝到长的,按位置覆盖
	copy(s3, s4)
	fmt.Println(s3)
	//修改4s的值
	s4[4] = 4000
	fmt.Println(s4)
	fmt.Println(s3) //修改s4的值后,s3的值不会发生变化(两个切片没有联系)
}
