package main

import (
	"fmt"
)

func main() {
	num1 := 0
	num2 := 0
	oper := ""

	fmt.Println("Please enter a number: ")
	fmt.Scanln(&num1)
	fmt.Println("Please enter another number: ")
	fmt.Scanln(&num2)
	fmt.Println("Please enter the operation: +, -, *, /")
	fmt.Scanln(&oper)

	switch oper{
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	}
}