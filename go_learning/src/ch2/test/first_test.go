package test

import "testing"

// 文件必须以 test.go 结尾
// 测试方法 必须以 Test 开头
func TestFirst(t *testing.T){

	t.Log(" TestFirst ")
}
