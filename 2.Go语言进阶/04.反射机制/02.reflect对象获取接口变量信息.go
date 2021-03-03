package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名:%s, 性别:%s, 年龄:%d", p.Name, p.Sex, p.Age)
}

func main() {
	//1.从实例到Value ValueOf()
	//2.从实例到Type TypeOf()
	//3.从Type到Value
	//	Type里只有类型信息,所以从Type无法获得实例Value,但可以通过type构建一个新的实例value.
	//	New(typ Type) Value 方法可以创建一个typ的指针类型
	//	Zero(typ Type) Value 方法可以返回一个typ类型的零值
	//	如果知道一个类型值的指针存放地址,是可以根据type和地址恢复value的
	//		func NewAt(typ Type, p unsafe.Pointer) Value
	//4.从Value到Type func (v Value) Type()
	//5.从Value到实例
	//	func(v value) Interface()
	//	func(v value) Bool/Float/Int/Uint
	//6.从Value指针到值,一个指针类型的Value获得值类型Value有两种方法
	//	func(v Value) Elem() //如果 v 类型是接口，则 Elem() 返回接口绑定的实例的 Value，如采 v 类型是指针，则返回指针值的 Value，否则引起 panic
	//	func Indirect(v Value) //如果 v 是指针，则返回指针值的 Value，否则返回 v 自身，该函数不会引起 panic
	//7.Type指针和值的相互转换
	//	t.Elem() 指针类型Type到值类型Type
	//	func PtrTo(t Type) 值类型Type到指针类型Type
	//8.Value值的可修改性
	//	func(v Value) CanSet() 判断是否可以修改
	//	func(v Value) Set() 修改

	//已知类型转换为其对应的类型
	var num float64 = 100.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	//转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
	//转换的时候，要区分是指针还是指
	//也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)

	p1 := Person{
		Name: "小明",
		Age:  18,
		Sex:  "男",
	}
	DoFiledAndMethod(p1)
}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get type is: ", getType.Name()) //Person
	fmt.Println("get kind is: ", getType.Kind()) //struct

	getValue := reflect.ValueOf(input)
	fmt.Println("get value is: ", getValue) //{小明 18 男}
	//获取方法字段
	//1. 先获取interface的reflect.Type,然后通过NumField进行遍历
	//2. 再通过reflect.Type的Field获取其Field
	//3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称:%s, 字段类型:%s, 字段值:%v\n", field.Name, field.Type, value)
	}

	//通过反射,操作方法
	// 1.先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	// 2.再通过reflect.Type的Method获取其Method
	for i := 0; i < getType.NumMethod(); i++ {
		methodName := getType.Method(i)
		fmt.Printf("方法名:%s, 方法类型:%v\n", methodName.Name, methodName.Type)
	}
}
