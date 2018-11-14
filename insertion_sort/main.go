package main

import "fmt"

func main() {

	arr := []int{10000, 500, 2, 4, 6, 9, 0}
	// res := insertionSort(arr)
	// fmt.Println(res)

	fmt.Println(insertSortAndFind(arr, 200))
}

// func insertionSort(A []int) []int {
// 	for j := 1; j < len(A); j++ {
// 		key := A[j]
// 		i := j - 1

// 		// nonincreasing order for i >= 0 && A[i] < key {
// 		for i >= 0 && A[i] > key {
// 			A[i+1] = A[i]
// 			i = i - 1
// 		}
// 		A[i+1] = key
// 	}
// 	return A
// }

// A = (1,2,3,...,n)
// v
// v =A[i]
// else v = nil

func insertSortAndFind(A []int, v int) (int, bool, []int) {
	if A[0] == v {
		return 0, true, A
	}

	for j := 1; j < len(A); j++ {
		if A[j] == v {
			return j, true, A
		}
		key := A[j]
		i := j - 1

		// nonincreasing order for i >= 0 && A[i] < key {
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}

	return 0, false, A

}
