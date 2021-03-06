package main

import (
	"fmt"
	"os"
)

/*
应用程序入口 要求
1. 必须有 main 函数
2. 包名一定是main 但是不一定是文件夹的名字
3. 文件名可以任意 不一定非要是main.go
*/

/*
运行go 程序

go run xxx.go
*/

// 无法通过 main 函数获取命令行参数
// 只能通过 os.Args 获取
func main() {

	if len(os.Args) > 1 {
		fmt.Println("Hello world", os.Args[1])
	}
	fmt.Println(os.Args)
}

/**

一般来说，需要为go添加三个主要的环境变量，分别是GOROOT, GOPATH以及go命令的路径，
通过安装包点击安装的go的位置默认为/usr/local/go，
因此需要将/usr/local/go/bin添加到PATH里，而GOROOT便是安装目录/usr/local/go，
GOPATH一般是自定义目录，用于存放你将要源码位置，如$HOME/go等，
修改环境变量可通过修改当前用户目录下的配置文件来修改：
*/
