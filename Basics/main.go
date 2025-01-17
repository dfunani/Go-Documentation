package main

import (
	"fmt"
	"slices"
)

type newStruct struct {
	age int
}

var anonStruct struct {
	age int
}

// anonStruct.age = 5

var anonStruct1 = struct {
	age int
}{
	age: 1,
}

func structs() {
	s := newStruct{
		7,
	}
	anonStruct.age = 6

	fmt.Println(s, anonStruct, anonStruct1)

}

func otherBuiltins() {
	var tempMap = make(map[string]int, 5)

	clear(tempMap)
	delete(tempMap, "lk")
	println(tempMap["k"])
	v, ok := tempMap["j"]
	if ok {
		println(v, tempMap["k"])
	} else {
		println("Nothing")
	}
}

func builtins(l int) {
	var tempSlice = make([]int, 5, 10)
	copy(tempSlice, tempSlice[:3])
	var tempSlice1 []int
	var tempSlice2 = []int{}
	var slice = []int{1, 2, 3}
	var arr = [11]int{1, 2, 3, 10: 100}
	slice = append(slice, 5)
	clear(slice)
	fmt.Println(slice, arr, tempSlice, tempSlice, tempSlice1)
	for i := range l {
		tempSlice1 = append(tempSlice1, i)
	}
	var tempSlice3 = tempSlice1[:3]
	fmt.Println(tempSlice1, tempSlice2, tempSlice3)
	tempSlice3 = append(tempSlice3, 9)
	fmt.Println(tempSlice1, tempSlice2, tempSlice3)

	fmt.Println(cap(tempSlice), cap(tempSlice1), cap(tempSlice2), cap(tempSlice3))
}

func main() {
	structs()
	otherBuiltins()
	const str = `Hello 'Again'`
	var defaultStr string

	var ints = 1
	var defaultInts int64 // uint8 == byte

	var floats float64 = 1
	var defaultFloats float64

	var char = 'a'
	var defaultChar rune

	var toggle = true
	var defaultToggle bool

	var arr = []int{1, 2, 3, 10: 100}
	var arr1 = [3]int{3, 2, 1}
	var arr2 = [3]int{1, 2, 3}
	var arr3 [1][3]int
	arr3[0][0] = 1
	arr3[0][1] = 1
	arr3[0][2] = 1
	var defaultArr [3]int

	var slice []int
	// slices.Sort(arr)
	slices.Reverse(arr)
	builtins(5)

	fmt.Println(float64(defaultFloats + floats))
	fmt.Println(arr, arr1, arr2, defaultArr)
	fmt.Println(arr3, slice)
	fmt.Println("Hello,", str, defaultStr, ints, defaultInts, floats, defaultFloats, string(char), defaultChar, toggle, defaultToggle)
}
