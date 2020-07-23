package main

import "fmt"

func main() {
	//定义二维数组
	var arr [2][2]int
	fmt.Println(arr)

	//赋值
	arr[0] = [2]int{1, 2}
	fmt.Println(arr)

	//定义同时赋值
	var arr2 = [...][2]string{{"a", "b"}, {"c", "c"}, {"e", "f"}}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("for-第%d行:", i)
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr2 {
		fmt.Printf("for-range-第%d行:", i)
		for _, v2 := range v {
			fmt.Print(v2)
		}
		fmt.Println()
	}
}
