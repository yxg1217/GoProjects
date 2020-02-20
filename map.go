package main

import (
	"fmt"
)

func main() {
	// Map: key-value
	//3 ways to create map
	var map1 map[string]float32 //one, map和slice是引用类型，初始值为nil，空slice能使用但空map不能使用，需要初始化
	var map2 = make(map[string]float32)//two
	var map3 = map[string]float32{"GO":98,"Python":90,"JavaScript":100}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1==nil)// 判断map是否为空
	fmt.Println(map2==nil)// 直接用make创建所以不为空
	fmt.Println(map3==nil)// 判断map是否为空

}