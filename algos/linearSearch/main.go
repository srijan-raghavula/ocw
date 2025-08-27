package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 8, 44, 9, 2}
	index := linearSearch(arr, 8)
	fmt.Println(arr)
	fmt.Println(index)
}

func linearSearch(arr []int, num int) int {
	for i, k := range arr {
		if k == num {
			return i
		}
	}
	return -1
}
