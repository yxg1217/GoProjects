package main

import (
	"fmt"
)

func main() {
	a := 1
	a++ 
	//++ --不能作为表达式，不能写作a:=a++，只能单独一行，++只能在变量右边
	var p *int = &a
	fmt.Println(*p)
	fmt.Println(p)
}