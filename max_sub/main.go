package main

import "fmt"

func main() {
	A := []int{-2, -1}
	n := len(A)
	fmt.Println(findMaxSub(A, 0, n-1))
}

func findMaxCrossingSub(A []int, low int, mid int, high int) (int, int, int) {
	var leftSum, rightSum, maxLeft, maxRight int
	sum := 0

	for i := mid; i >= low; i-- {
		sum = sum + A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum = sum + A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

func findMaxSub(A []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, A[low]
	}
	if len(A) < 3 {
		if A[low] > A[high] {
			return low, high, A[low]
		} else {
			return low, high, A[high]
		}
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaxSub(A, low, mid)
	rightLow, rightHigh, rightSum := findMaxSub(A, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSub(A, low, mid, high)
	//fmt.Println(mid, leftLow, leftHigh, leftSum, rightLow, rightHigh, rightSum, crossLow, crossHigh, crossSum)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}

	return crossLow, crossHigh, crossSum
}
