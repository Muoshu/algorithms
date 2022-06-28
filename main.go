package main

import (
	"fmt"
	"time"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func main() {
	//var x = []int{53, 17, 78, 9, 45, 65, 87, 32}
	//fmt.Println("original:", x)
	//
	//sort.HeapSort(x)
	start := time.Now().Nanosecond()
	fmt.Println(f(20))
	end := time.Now().Nanosecond()
	fmt.Println(end - start)

}

func f(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return f(n-1) + f(n-2)

}
