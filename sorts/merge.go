package sorts

import (
	"sync"
)

// MergeSort is an implementation of a mergesort (noshit?)
// <3 slices
// Time complexity: O(n * log n)
// Space complexity: O(log n + n) // <- Not really sure about this one
func MergeSort(input []int) {
	helper := make([]int, len(input))

	mergeSort(input, helper)
}

func mergeSort(input, helper []int) {
	if len(input) <= 1 {
		return
	}

	mergeSort(input[:len(input)/2], helper[:len(input)/2])
	mergeSort(input[len(input)/2:], helper[len(input)/2:])
	merge(input, helper)
}

// ParallelMergeSort is a mergeSort but done in parallel
func ParallelMergeSort(input []int) {
	helper := make([]int, len(input))

	parallelMergeSort(input, helper)
}

func parallelMergeSort(input, helper []int) {
	if len(input) <= 1 {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		mergeSort(input[:len(input)/2], helper[:len(input)/2])
		wg.Done()
	}()
	go func() {
		mergeSort(input[len(input)/2:], helper[len(input)/2:])
		wg.Done()
	}()

	wg.Wait()

	merge(input, helper)
}

func merge(input, helper []int) {
	copy(helper, input)

	current, left, right := 0, 0, len(input)/2

	for left < len(input)/2 && right < len(input) {
		if helper[left] < helper[right] {
			input[current] = helper[left]
			left++
		} else {
			input[current] = helper[right]
			right++
		}

		current++
	}

	for left < len(input)/2 {
		input[current] = helper[left]
		current++
		left++
	}

	for right < len(input) {
		input[current] = helper[right]
		current++
		right++
	}
}
