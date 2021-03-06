package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {

	// 声明变量
	// 方式一 直接声明类型

	var a int = 1
	var b int = 1

	// 方式二 方式1的简化方式
	// var (
	// 	a int = 1
	// 	b int = 1
	// )

	// 方式三 类型推断

	// a:=1
	// b:=1

	t.Log(a)

	for i := 0; i < 5; i++ {
		t.Log(b)

		tmp := a + b
		b = tmp
		a = b
	}
}

func TestNine(t *testing.T) {
	var a int = 1
	var b int = 2

	b = 1000

	fmt.Println("以上是我认为的声明变量的简洁方式", b)

	for i := 0; i < 9; i++ {
		fmt.Println(i, a)
		a += 1
	}
}
