
package main

import "fmt"

func main(){
	fmt.Println("Hello")

	var a int = 1;
	var b int = 1;

	// 方式二

	// a:=1
	// b:=1

	// 方式三
	// var (
	// 	a int = 1
	// 	b int = 1
	// )

	// 1 1 2 3 5 8 13
	fmt.Println(a)

	for i := 0; i < 5; i ++{
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + b

	}
}