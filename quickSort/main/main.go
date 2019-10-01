package main

import "fmt"

func partition(slice []int, low, high int) int {
	i := low - 1
	pivot := slice[high]
	for j := low; j < high; j++ {
		if slice[j] < pivot {
			i = i + 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSort(slice []int, low, high int) {
	if low < high {
		pi := partition(slice, low, high)
		quickSort(slice, low, pi-1)
		quickSort(slice, pi+1, high)
	}
	return
}

func main() {
	arr := [...]int{111111, 2, 0, 10, 34353, 32423, 223, 455, 11, 34}
	slice := make([]int, 10)
	slice = arr[:]
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}
