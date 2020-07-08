package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 60	// 00111100
	var b = 13	// 00001101
	//与
	var c = a&b	// 00001100
	fmt.Println(c, strconv.FormatInt(int64(c), 2))
	//或
	var d = a|b	// 00111101
	fmt.Println(d, strconv.FormatInt(int64(d), 2))
	//亦或
	var e = a^b	// 00110001
	fmt.Println(e, strconv.FormatInt(int64(e), 2))
	var f = a>>1
	fmt.Println(f, strconv.FormatInt(int64(f), 2))
	var g = f<<1
	fmt.Println(g, strconv.FormatInt(int64(g), 2))
}