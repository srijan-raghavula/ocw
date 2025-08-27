package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 9, 3, 1}
	fmt.Println(arr)
	arr = mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	left := mergeSort(arr[:n/2])
	right := mergeSort(arr[n/2:])
	return merge(left, right)
}

func merge(one, two []int) []int {
	lOne, lTwo := len(one), len(two)
	sortedArr := make([]int, 0)
	i, j := 0, 0
	for i < lOne && j < lTwo {
		if one[i] <= two[j] {
			sortedArr = append(sortedArr, one[i])
			i++
		} else {
			sortedArr = append(sortedArr, two[j])
			j++
		}
	}
	if i < lOne {
		sortedArr = append(sortedArr, one[i:]...)
	}
	if j < lTwo {
		sortedArr = append(sortedArr, two[j:]...)
	}
	return sortedArr
}
