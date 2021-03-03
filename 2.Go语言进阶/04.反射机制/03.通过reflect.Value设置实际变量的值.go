package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Name  string
	Color string
}

func main() {
	//修改值
	var num float64 = 1.2345
	fmt.Println("num的值:", num)

	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem() //如果非指针,会直接panic

	fmt.Println("类型:", newValue.Type())
	fmt.Println("是否可以修改:", newValue.CanSet())

	newValue.SetFloat(4.567)
	fmt.Println(num)
	fmt.Println("===========================")

	//修改结构体中的值
	bmw := Car{
		Name:  "宝马",
		Color: "白色",
	}
	fmt.Println(bmw)
	bmw2 := &bmw
	fmt.Printf("%T\n", bmw)
	fmt.Printf("%T\n", bmw2)
	fmt.Println(bmw.Name, bmw2.Name)

	car := reflect.ValueOf(&bmw)
	if car.Kind() != reflect.Ptr {
		fmt.Println("car不是指针类型,无法修改值")
		return
	}
	fmt.Println("是否可修改:", car.Elem().CanSet())
	car1 := car.Elem()
	color := car1.FieldByName("Color")
	color.SetString("红色") //修改颜色成功
	fmt.Println(bmw)
}
