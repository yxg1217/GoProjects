package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	[100, 999]
	Un nombre narcissique est un entier naturel n non nul 
	qui est égal à la somme des puissances p-ièmes de ses chiffres 
	en base dix, où p désigne le nombre de chiffres de n
	一水仙花数所谓的水仙花数是指：
	一个n 位数( n≥3 )，它的每个位上的数字的n 次幂之和等于它本身。
	 例如153，370，371，407等都是水仙花数，
	 就拿153来说， 153=1*1*1 + 5*5*5 + 3*3*3 

	285:
	285 / 100 = 2
	285 % 10 = 5
	285 --> 28 %10 = 8
	*/

	for i := 100; i < 1000; i++{
		x := i / 100 //百
		y := i /10 % 10 //十
		z := i % 10 //个

		if math.Pow(float64(x),3) + math.Pow(float64(y),3) + math.Pow(float64(z),3) == float64(i) {
			fmt.Println(i)
			// 153, 370, 371 ,407
		}
	}

	fmt.Println("--------------------------------")

	/*
	hundred: 1 - 9
	ten: 0 - 9
	one: 0 - 9
	*/
	for a := 1; a < 10; a++ { //hundred
		for b := 0; b < 10; b++ { //ten
			for c := 0; c < 10; c++ { //one
				n := a * 100 + b * 10 + c * 1
				if a*a*a + b*b*b + c*c*c == n {
					fmt.Println(n)
				}
			}
		}
	}
}
