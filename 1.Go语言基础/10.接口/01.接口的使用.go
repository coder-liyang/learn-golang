package main

import "fmt"

//在go语言中,接口是一组方法签名,当类型为接口中的所有方法被定义时,成为实现接口

type PC interface {
	MainBoard() string
	CPU() string
	Mem() string
	Disk() string
}

type Usb interface {
	GetUsbVersion() string
}

type Mac struct {
	Name  string
	Color string
	Price float64
}

func (m *Mac) MainBoard() string {
	return "MainBoard"
}

func (m *Mac) CPU() string {
	return "CPU"
}
func (m *Mac) Mem() string {
	return "Mem"
}
func (m *Mac) Disk() string {
	return "Disk"
}

//interface可以被任意的对象实现
//一个对象可以实现任意多个interface的interface
//任意的类型都实现了空interface,也就是包含0个method的
func (m *Mac) GetUsbVersion() string {
	return "3.0"
}

func main() {
	mac := Mac{
		Name:  "Mac",
		Color: "Write",
		Price: 10000,
	}
	fmt.Println(mac.Name, mac.CPU(), mac.Disk())

	var mac2 = new(Mac)
	mac2.Name = "MacAri"

	fmt.Println(mac2.Name, mac2.Mem(), mac2.GetUsbVersion())
}
