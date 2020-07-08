package main

import (
	"fmt"
	"strconv"
)

func main() {
	//布尔型
	var a = true
	var b = false
	fmt.Println(a, b)
	//数值型
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1
	var ui8 uint8 = 5
	var i = 1
	var ii = 1
	//fmt.Println(i8 + i16) //mismatched types int8 and int16
	fmt.Println(int16(i8) + i16)
	fmt.Printf("%T, %T, %T, %T, %T, %T \r\n", i8, i16, i32, i64, i, ii)
	var bb byte = 123
	fmt.Printf("%d, %T\r\n", bb+ui8, bb)

	//浮点型
	var f1 float32 = 1.1
	var f2 = 2.2
	var f3 = 3.3
	var c1 complex64 = 1.1
	var c2 complex128 = 1.1
	fmt.Printf("%T, %T, %T,%T,%T\r\n", f1, f2, f3, c1, c2)

	//字符串型
	var str1 = "abcdefg"
	var str2 = "abcdefg"
	var str3 string
	fmt.Println(str1, str2, str3)
	var str4 = strconv.FormatFloat(f3, 'f', 2, 64)

	//数据类型转换
	fmt.Printf("%T, %T, %T, %T, %s\r\n", f1, float64(f1), strconv.Itoa(i), str4, str4)
}
