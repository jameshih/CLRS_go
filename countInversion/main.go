package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 20, 6, 4, 5}
	sorted, inversion := mergeSortnCountInvert(arr)
	fmt.Println(sorted, "| number of inversions:", inversion)
}

func merge(left, right []int) ([]int, int) {

	size, i, j, inversion := len(left)+len(right), 0, 0, 0
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
			inversion += (len(left) - i)
		}
	}
	return A, inversion
}

func mergeSortnCountInvert(A []int) ([]int, int) {

	if len(A) < 2 {
		return A, 0
	}
	mid := (len(A)) / 2
	a, ai := mergeSortnCountInvert(A[:mid])
	b, bi := mergeSortnCountInvert(A[mid:])
	c, ci := merge(a, b)
	return c, (ai + bi + ci)
}
