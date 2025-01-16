package main

import (
	"fmt"
)

func main() {
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

	fmt.Println(float64(defaultFloats + floats))
	fmt.Println("Hello,", str, defaultStr, ints, defaultInts, floats, defaultFloats, string(char), defaultChar, toggle, defaultToggle)
}
