package constant

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	ThuesDay
	Wednesday
	ThursDay
)

func TestConst(t *testing.T) {
	fmt.Println(Monday, Wednesday)
}
