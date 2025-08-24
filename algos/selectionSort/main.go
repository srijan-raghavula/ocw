package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 2, 4, 9, 3}
	selectionSort(arr, len(arr)-1)
	fmt.Println(arr)
}

func selectionSort(arr []int, i int) {
	if i > 0 {
		j := prefixMax(arr, i)
		swapElements(arr, i, j)
		selectionSort(arr, i-1)
	}
}

func prefixMax(arr []int, i int) int {
	if i > 0 {
		j := prefixMax(arr, i-1)
		if arr[i] < arr[j] {
			return j
		}
	}
	return i
}

func swapElements(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
