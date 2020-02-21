package main

import (
	"fmt"
	"sort"
)

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 7, 10, 84}
	//要求对切片进行排序
	// 1.冒泡排序
	// 2.也可以使用系统提供的方法 sort
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}