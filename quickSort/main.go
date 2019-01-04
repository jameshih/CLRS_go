package main

import "fmt"

func main() {
	arr := []int{10, 2, 35, 4, 111, 20, 3, 10, 1, 200, 4, 99, 3, 0}
	res, _, _ := quickSort(arr, 0, len(arr)-1)

	fmt.Println(res)
}

func quickSort(A []int, left, right int) ([]int, int, int) {
	pivot, partitionIndex := 0, 0

	if left < right {
		pivot = right
		partitionIndex = partition(A, pivot, left, right)

		quickSort(A, left, partitionIndex-1)
		quickSort(A, partitionIndex+1, right)
	}
	return A, left, right
}

func partition(A []int, pivot, left, right int) int {
	pivotValue := A[pivot]
	partitionIndex := left

	for i := left; i < right; i++ {
		if A[i] < pivotValue {
			swap(A, i, partitionIndex)
			partitionIndex++
		}
	}

	swap(A, right, partitionIndex)
	return partitionIndex
}

func swap(A []int, firstIndex, secondIndex int) {
	temp := A[firstIndex]
	A[firstIndex] = A[secondIndex]
	A[secondIndex] = temp
}
