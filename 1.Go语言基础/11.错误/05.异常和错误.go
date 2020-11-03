package main

import (
	"errors"
	"fmt"
)

//失败处理的正确姿势
//1.失败原因只有一个时,不使用error
//该函数只有一个错误原因
func CheckHostType(host_type string) error {
	switch host_type {
	case "virtual_machine":
		return nil
	case "bare_metal":
		return nil
	}
	return errors.New("CheckHostType ERROR:" + host_type)
}

//优化后
func IsValidHostType(hostType string) bool {
	return hostType == "virtual_machine" || hostType == "bare_metal"
}

//2.没有失败时,不使用error
//该函数没有失败的可能
func setTenantId(ii int64) error {
	iii := ii
	fmt.Println(iii)
	return nil
}

//优化后
func setTenantId2(ii int64) {
	iii := ii
	fmt.Println(iii)
}

//3.error应该放在返回值类型的最后[bool作为返回值也一样]
func say(str string) (string, error) {
	return str, nil
}

//4.错误值统一定义
//5.错误层层传递时,每层都加日志
//6.错误处理使用defer
//7.当尝试几次可以避免失败时,不要立即返回错误
//8.当上层函数不关心错误时,最好不返回error
//9.当错误发生时,不忽略有用的返回值

//异常处理的正确姿势
//1.在程序开发阶段,坚持速错
//2.在程序部署后,应恢复异常避免程序终止
//3.对于不应该出现的分支,使用异常处理
//4.针对入参不应该有问题的函数,使用panic设计
func main() {

}
