package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Date(2020, time.Month(time.February), 1, 1, 1, 1, 1, time.Local))

	fmt.Println("日期转文本:", time.Now().Format("2006年01月02日 15时04分05秒"))
	parseTime, _ := time.Parse("2006年01月02日 15时04分05秒", "2021年01月02日 13时49分39秒")
	fmt.Println("文本转日期", parseTime)

	fmt.Println("当前时间戳(秒):", time.Now().Unix())
	fmt.Println("当前时间戳(纳秒):", time.Now().UnixNano())

	fmt.Println(time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println("现在时间是:" + time.Now().Format("2006年01月02日 15时04分05秒") + ",程序马上睡5秒")
	time.Sleep(time.Second * 5)
	fmt.Println("现在时间是:" + time.Now().Format("2006年01月02日 15时04分05秒") + ",程序睡醒了")

}
