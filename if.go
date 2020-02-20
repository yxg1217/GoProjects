package main

import (
	"fmt"
)

func main() {
	a := 10
	if a := 1; a > 0 {
		//if语句的判断不需要括号，前面可以加初始 a：= 1， 但只能在if语句块里使用，第二个print就显示不出if后的a
		fmt.Println(a)
	}
	fmt.Println(a)
}