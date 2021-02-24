package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	scanDir("./08.遍历文件夹", 0)
}

func scanDir(path string, level int) {
	level++
	prefix := "|-"
	for i := 0; i <= level; i++ {
		prefix += "--"
	}

	fileInfo, _ := ioutil.ReadDir(path)
	for _, v := range fileInfo {
		if v.IsDir() {
			fmt.Println(prefix+"dir:", path+"/"+v.Name())
			scanDir(path+"/"+v.Name(), level)
		} else {
			fmt.Println(prefix+"file:", path+"/"+v.Name())
		}
	}
}

func scanDir2(path string) {

}
