package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	// 数组的声明
	var arr [3]int = [3]int{1, 2, 3}

	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}

	t.Log("arr", arr)
	t.Log(arr1[0], arr1[1], arr1[3])
	t.Log("arr3", len(arr3))
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}

	// 普通遍历
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	// for each
	for idx, value := range arr3 {
		t.Log("arr3", idx, value)
	}
}

func TestArraySection(t *testing.T) {
	// 数组切片
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	arr3_sestion := arr3[1:3]
	arr3_sestion1 := arr3[:3]           //前三个
	arr3_sestion2 := arr3[len(arr3)-3:] // 后三个
	t.Log("arr3_sestion", arr3_sestion)
	t.Log("arr3_sestion1", arr3_sestion1)
	t.Log("arr3_sestion1", arr3_sestion2)
}

func TestSliceInit(t *testing.T) {

	var s0 []int
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	var s1 []int = []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	s1 = append(s1, 5)
	t.Log(len(s1), cap(s1)) // 此时已经扩容

	// 创建一个切片 长度是3 容量是 5
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	var s []int

	for i := 0; i < 20; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {

	var s []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

	Q0 := s[3:6]
	t.Log("Q0", Q0, len(Q0), cap(Q0)) // Q0, [D E F], 3, 9

	Q0[0] = "<<<<"
	t.Log("s", s)

}
