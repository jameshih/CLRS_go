package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 41, 52, 26, 38, 57, 9, 49}
	fmt.Println(mergeSort(arr))
}

func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	A := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			A[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			A[k] = left[i]
			i++
		} else if left[i] < right[j] {
			A[k] = left[i]
			i++
		} else {
			A[k] = right[j]
			j++
		}
	}
	return A
}

func mergeSort(A []int) []int {

	if len(A) < 2 {
		return A
	}
	mid := (len(A)) / 2
	return merge(mergeSort(A[:mid]), mergeSort(A[mid:]))
}
