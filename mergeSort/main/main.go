package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func mergeSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}

	mid := len(slice) / 2

	left := make([]int, mid)
	right := make([]int, len(slice)-mid)

	left = slice[:mid]
	right = slice[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func main() {
	arr := []int{234812348, 20, 0, 231413, 232, 22, 3123, 434, 233, 2, 12}
	slice := mergeSort(arr)
	fmt.Println(slice, arr)
}
