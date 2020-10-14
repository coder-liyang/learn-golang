package main

import "fmt"

type computer struct {
	cpu  string
	mem  int
	disk int
	usb  interface{}
}

func main() {
	pc1 := computer{
		cpu:  "i5",
		mem:  4000000,
		disk: 500000000,
		usb:  "keyboard",
	}
	printPc(pc1)

	var pc2 = new(computer)

	pc2.cpu = "i7"
	pc2.mem = 8000000
	pc2.disk = 500000000
	pc2.usb = "mouse"
	printPc(*pc2)

}

func printPc(pc computer) {
	fmt.Println(pc)
}
