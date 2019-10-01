package main

import "fmt"

func selectionSort(slice []int) {
	for i, _ := range slice {
		min_index := i
		for j := i + 1; j < len(slice); j++ {
			if slice[min_index] > slice[j] {
				min_index = j
			}
			slice[i], slice[min_index] = slice[min_index], slice[i]
		}
	}
	return
}

func main() {
	arr := [...]int{2, 11144, 325, 34, 1, 44325235, 324200}
	slice := make([]int, 1)
	slice = arr[:]
	selectionSort(slice)
	fmt.Println(slice)
}
