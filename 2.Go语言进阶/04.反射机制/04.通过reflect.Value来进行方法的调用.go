package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) Say() {
	fmt.Println("Say()方法\t姓名:", s.Name, ";", "年龄:", s.Age)
}

func (s Student) Addition(a, b int) {
	fmt.Println(a, "+", b, "=", a+b)
}

func main() {
	s1 := Student{"张三", 10}
	s1.Say()

	s2 := reflect.ValueOf(&s1)
	fmt.Println("CanSet:", s2.Elem().CanSet())
	s2.Elem().FieldByName("Age").SetInt(18)
	fmt.Println(s1)
	method := s2.Elem().MethodByName("Say")
	//没有参数
	method.Call(nil)                  //直接写nil
	args1 := make([]reflect.Value, 0) //或者创建一个空的切片也可以
	method.Call(args1)

	//有参数
	method2 := s2.Elem().MethodByName("Addition")
	args2 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200)}
	method2.Call(args2)

}
