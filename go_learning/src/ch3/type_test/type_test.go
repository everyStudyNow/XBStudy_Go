package type_test

import (
	"fmt"
	"testing"
)

/*
go 语言中基本数据类型

1. bool
2. string
3. int int8 int16 int32 int64 (为了适配不同的运行环境)
4. uint 无符号的int
5. byte int8 的别名
6. rune int32 的别名
7. float32 float64
8. complex64 complex128 复数
*/

/*
不支持隐式类型转换 就是小类型到大类型的转换 比如 int32 可以赋值给 int64
*/

func TestType(t *testing.T) {
	var a int
	var b int64 = 2
	a = int(b)

	fmt.Println(a)
}

/*
类型的预定义值

math.MaxInt64
math.MaxFloat64
math.MaxUnit32
*/

/*
指针类型

不支持指针的运算
string 是值类型 默认值不是nil 是一个空字符串
*/

func TestPoint(t *testing.T) {
	var a int = 1
	var aPtr = &a //取出 a的指针的地址
	fmt.Println(a, aPtr)
	// aPtr = aPtr + 1 //invalid operation: aPtr + 1 不支持指针运算
}

func TestString(t *testing.T) {
	var s string
	fmt.Println("********")
	fmt.Println(len(s))
}
