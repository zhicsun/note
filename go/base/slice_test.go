package base

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	var a []int
	fmt.Println(a == nil) // true

	a = append(a, 1)
	fmt.Println(a) // 1
}

func TestAddCap(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := slice[2:5]     // 2, 3, 4
	fmt.Println(cap(s1)) // 8

	s2 := s1[2:6:7]      // 4, 5, 6, 7
	fmt.Println(cap(s2)) // 5

	s2 = append(s2, 100) // 4, 5, 6, 7, 100
	s2 = append(s2, 200) // 4, 5, 6, 7, 100, 200

	s1[2] = 20

	fmt.Println(s1)    // 2, 3, 20
	fmt.Println(s2)    // 4, 5, 6, 7, 100, 200
	fmt.Println(slice) // 0, 1, 2, 3, 20, 5, 6, 7, 100, 9
}

func TestGroup(t *testing.T) {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
	// newcap=5 ptrSize=8 smallSizeDiv=1024 class_to_size[size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv]]
	// capmem = roundupsize(uintptr(newcap) * ptrSize)
	// newcap = int(capmem / ptrSize)
}

func f(s []int) {
	for i := range s {
		s[i] += 1
	}
}

func myAppend(s []int) []int {
	s = append(s, 100)
	s = append(s, 200)
	return s
}

func myAppendPtr(s *[]int) {
	*s = append(*s, 100)
	return
}

func TestFunc(t *testing.T) {
	s := []int{1, 1, 1}
	f(s) // 修改原切片

	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newS := myAppend(s[7:8:9]) // 发生扩容不影响原切片
	fmt.Println(s)
	fmt.Println(newS)

	s = newS
	myAppendPtr(&s) // 修改原切片
	fmt.Println(s)
}
