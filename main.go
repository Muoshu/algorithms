package main

import (
	"algorithms/design"
	"fmt"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func main() {
	var x = []int{1, 1, 0, 1, 1, 1, 3, 2}

	fmt.Println(design.FindMinMax2(x))

}
