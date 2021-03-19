package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a1 int8
	a2 int16
} //4
type B struct {
	b1 int16
	b2 int8
	b3 int //8
} //16
type C struct {
	c1 int8
	c2 float32 //4
	c3 int
} //16
type D struct {
	d1 float32
	d2 int
	d3 int8
} //24

func main() {
	var a = A{}
	var b = B{}
	var c = C{}
	var d = D{}
	var e int
	var f float32
	var g int32
	var h int64
	fmt.Printf("size of A = %d\n", unsafe.Sizeof(a))
	fmt.Printf("size of B = %d\n", unsafe.Sizeof(b))
	fmt.Printf("size of C = %d\n", unsafe.Sizeof(c))
	fmt.Printf("size of D = %d\n", unsafe.Sizeof(d))
	fmt.Printf("size of int = %d\n", unsafe.Sizeof(e))
	fmt.Printf("size of float32 = %d\n", unsafe.Sizeof(f))
	fmt.Printf("size of int32 = %d\n", unsafe.Sizeof(g))
	fmt.Printf("size of int64 = %d\n", unsafe.Sizeof(h))

	clickNum := 10
	clickNumPos := clickNum - 1
	ClickItem := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11}
	realClickItemLen := len(ClickItem)
	tmpClickItem := make([]int, clickNum, clickNum)
	if realClickItemLen == clickNum {
		tmpClickItem = ClickItem
	} else if realClickItemLen < clickNum {
		for i := realClickItemLen - 1; i >= 0; i-- {
			tmpClickItem[clickNumPos] = ClickItem[i]
			clickNumPos--
		}
	}
	fmt.Println(ClickItem)
	fmt.Println(tmpClickItem)

	//var vecA []int = []int{1, 2, 3}
	var vecA []int
	var mapB map[int]int
	mapB = make(map[int]int)
	vecA = append(vecA, 1)
	mapB[1] = 1
	fmt.Println(mapB)
	fmt.Println(vecA)
}
