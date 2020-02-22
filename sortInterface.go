package main

import (
	"fmt"
	"sort"
	"math/rand"
)

//1.申明hero结构体
type Hero struct{
	Name string
	Age int
}

//2.申明hero结构体的切片类型，使其可以存放多个hero

type HeroSlice []Hero

//3. 实现官网https://golang.org/pkg/sort/#Interface 的三个方法
func (hs HeroSlice) Len() int {
	return len(hs)
}
//下表i<J为真则按从小到大升序排列，i>j 则反之
func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i,j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
} //三个方法都实现，意味着实现了Interface接口，此时我们可以调用sort包里的Sort方法

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 7, 10, 84}
	//要求对切片进行排序
	// 1.冒泡排序
	// 2.也可以使用系统提供的方法 sort
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	/*
	Sort 接受一个interface，interface里有三个方法，因此要对一个结构体排序，需要实现interface里的三个方法
	*/

	//要求对结构体切片进行排序
	//1.冒泡排序
	//2.Sort方法

	//测试是否可以对结构体切片排序
	var heroes HeroSlice
	for i := 0 ; i < 10 ; i ++ {
		hero := Hero{
			Name : fmt.Sprintf("Hero%d", rand.Intn(100)), //Math package, return a random number between (0,n)
			Age : rand.Intn(100),
		}
		//append hero to the slice
		heroes = append(heroes, hero)
	}

	//Before quickSort
	for _ , v := range heroes {
		fmt.Println(v)
	}

	// Call sort.Sort
	sort.Sort(heroes)
	//After quickSort
	fmt.Println("------------After quickSort------------")
	for _ , v := range heroes {
		fmt.Println(v)
	}

}



