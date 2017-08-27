package sorts

import (
	"reflect"
	"testing"
)

func clone(in []int) []int {
	res := make([]int, len(in))
	copy(res, in)
	return res
}

func TestSorts(t *testing.T) {
	cases := []struct {
		Label       string
		Input       []int
		Expectation []int
	}{
		{
			"NominalCase",
			[]int{0, 9, 5, 43, 10, 8, 2},
			[]int{0, 2, 5, 8, 9, 10, 43},
		},
		{
			"EmptyArray",
			[]int{},
			[]int{},
		},
		{
			"EqualAndNegativesValues",
			[]int{-4, 0, 204, 500, 2, 3934, 8, 5, 2},
			[]int{-4, 0, 2, 2, 5, 8, 204, 500, 3934},
		},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			mergeOutput := clone(test.Input)
			MergeSort(mergeOutput)
			if !reflect.DeepEqual(mergeOutput, test.Expectation) {
				t.Error("Invalid result for MergeSort", mergeOutput, test.Expectation)
			}
		})
	}
}
