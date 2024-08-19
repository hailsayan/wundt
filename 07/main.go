package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	maxSum := arr[0]
	currentSum := arr[0]

	for i := 1; i < n; i++ {
		currentSum = max(arr[i], currentSum+arr[i])
		maxSum = max(maxSum, currentSum)
	}

	fmt.Println(maxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
