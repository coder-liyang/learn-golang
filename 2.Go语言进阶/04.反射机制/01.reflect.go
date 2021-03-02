package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64
	x = 100.32
	//通过反射获取变量的类型和值
	fmt.Println("类型:", reflect.TypeOf(x)) //TypeOf()得到真实的类型
	fmt.Println("值:", reflect.ValueOf(x)) //ValueOf()得到具体的值

	//根据反射的值,来判断类型和数值
	value := reflect.ValueOf(x)
	fmt.Println("Kind is float64", value.Kind() == reflect.Float64) //Kind()返回具体的信息
	fmt.Println("Kind is float32", value.Kind() == reflect.Float32)

	fmt.Println("type: ", value.Type())
	fmt.Println("value: ", value.Float())

	a := 100
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Type())
}
