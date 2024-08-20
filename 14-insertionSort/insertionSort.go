package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		p := i
		item := arr[p]

		for p > 0 && item < arr[p-1] {
			arr[p] = arr[p-1]
			p = p - 1
		}
		arr[p] = item
	}
	return arr
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sortedArr := insertionSort(arr)

	for i, v := range sortedArr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
