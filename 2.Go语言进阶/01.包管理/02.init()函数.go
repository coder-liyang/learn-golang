package main

//main和init为系统保留函数
//init会在main之前执行
//main和init不能有任何返回值,系统自动调用,不可被引用
//init函数可以有多个

//对同一个 go 文件的 init( ) 调用顺序是从上到下的。
//对同一个 package 中的不同文件，将文件名按字符串进行“从小到大”排序，之后顺序调用各文件中的init()函数。
//对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数。
//如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，一次执行对应的 init 方法。main 包总是被最后一个初始化，因为它总是依赖别的包
//一个包被其它多个包 import，但只能被初始化一次
import (
	"fmt"
	_ "learn-golang/2.Go语言进阶/01.包管理/pkg"
)

func main() {
	fmt.Println("This is main.main()")
}

func init() {
	fmt.Println("This is main.init() 1")
}
func init() {
	fmt.Println("This is main.init() 2")
}
