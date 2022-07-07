package main

import (
	"algorithms/lc"
	"fmt"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func main() {
	//var x = []int{8,2,3,5,7,6,11,14,1,9,13,10,4,12,15,16}
	//fmt.Println(design.SelectK(x,0,15,6))
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"

	fmt.Println(lc.ReplaceWords(dictionary, sentence))

}
