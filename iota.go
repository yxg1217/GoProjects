package main

import (
	"fmt"
)

const(
	B float64 = 1 << (iota * 10)
	KB
	MB
	TB
	PB
	EB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
}