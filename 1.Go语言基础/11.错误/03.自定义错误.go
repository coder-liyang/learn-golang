package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//创建错误可以用errors包下得New()函数,以及fmt包下的Errorf()函数

	area, err := circleArea2(10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(area)
	}

	area, err = circleArea3(-7)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Println("面积计算错误")
			fmt.Println(err.Error(), err.err, err.radius)
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(area)
	}
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//如果想格式化错误信息,用这个方法不太友好
		return 0, errors.New("半径必须>0")
	}
	return radius * radius * math.Pi, nil
}

func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		//使用fmt.Errorf可以更好的格式化错误信息
		return 0, fmt.Errorf("值:%0.2f错误", radius)
	}
	return radius * radius * math.Pi, nil
}

//以上两个如果要判断具体错误都只能判断错误文案,而文案很可能会更改,这样的判断就会有问题
//正确的方法应该是将错误实现为错误的struct类型
func circleArea3(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{fmt.Sprintf("值:%0.2f错误", radius), radius}
	}
	return radius * radius * math.Pi, nil
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius error: %.2f, %s", e.radius, e.err)
}
