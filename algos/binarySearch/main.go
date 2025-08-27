package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 3, 6, 7, 38}
	target := 7
	index := binarySearch(arr, target)
	fmt.Println(arr)
	fmt.Println(index)
}

func binarySearch(arr []int, target int) int {
	i, j := 0, len(arr)
	middle := (i + j) / 2
	for {
		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			middle += middle / 2
		} else {
			middle -= middle / 2
		}
	}
}
