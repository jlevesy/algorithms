package sorts

import (
	"fmt"
)

// QuickSort is an implementation of a QuickSort
func QuickSort(input []int) {
	fmt.Println("Quicksort with input", input)
	index := partition(input)

	left := input[:index]
	right := input[index:]

	if len(left) > 1 {
		QuickSort(left)
	}

	if len(right) > 1 {
		QuickSort(right)
	}
}

func partition(input []int) int {
	pivot := input[len(input)/2]
	left, right := 0, len(input)-1
	for left < right {
		for input[left] < pivot {
			left++
		}

		for input[right] > pivot {
			right--
		}

		input[left], input[right] = input[right], input[left]

		left++
		right--
	}

	return left
}
