package main

import (
	"algorithms/sort"
	"fmt"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func main() {
	var x = []int{53, 17, 78, 9, 45, 65, 87, 32}
	fmt.Println("original:", x)
	sort.HeapSort(x)
	fmt.Println(x)

}
