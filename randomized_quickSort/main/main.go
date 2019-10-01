package main

import (
	"fmt"
	"math/rand"
)

func partition(slice []int, start, stop int) int {
	pivot := start
	i := start + 1

	for j := start + 1; j < stop+1; j++ {
		if slice[j] <= slice[pivot] {
			slice[i], slice[j] = slice[j], slice[i]
			i = i + 1
		}
	}
	slice[pivot], slice[i-1] = slice[i-1], slice[pivot]
	pivot = i - 1
	return pivot
}

// Lomuto Partitioning implementation for randomized pivot sub routine
func partitionrand(slice []int, start, stop int) int {
	randpivot := rand.Intn(stop-start) + start //equal to radom range(start,stop)
	slice[start], slice[randpivot] = slice[randpivot], slice[start]
	return partition(slice, start, stop)
}

func quickSort(slice []int, start, stop int) {
	if start < stop {
		pivotIndex := partitionrand(slice, start, stop)

		quickSort(slice, start, pivotIndex-1)
		quickSort(slice, pivotIndex+1, stop)
	}
	return
}

func main() {
	arr := [...]int{3242, 1232, 5431, 135, 7782, 3432534, 333, 4, 22, 8, 243547, 35, 2}
	slice := make([]int, 1)
	slice = arr[:]
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}
