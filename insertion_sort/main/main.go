package main

import "fmt"

func insertionSort(slice []int) {
	for i, key := range slice {
		j := i - 1
		for j >= 0 && key < slice[j] {
			slice[j+1] = slice[j]
			j -= 1
			slice[j+1] = key
		}
	}
	return
}

func main() {
	arr := [...]int{44, 2, 111, 2, 32432432, 1, 3343, 34, 7}
	slice := make([]int, 1)
	slice = arr[:]
	insertionSort(slice)
	fmt.Println(slice)

}
