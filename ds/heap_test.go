package ds

import (
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	newCases := []struct {
		Label       string
		In          MaxIntHeap
		Expectation MaxIntHeap
	}{
		{
			"BaseCase",
			MaxIntHeap{9, 10, 49, 3, 593, 4},
			MaxIntHeap{593, 49, 10, 3, 9, 4},
		},
		{
			"WithEqualTerms",
			MaxIntHeap{9, 9, 10, 49, 3, 593, 4},
			MaxIntHeap{593, 49, 10, 9, 3, 9, 4},
		},
	}

	for _, test := range newCases {
		t.Run(test.Label, func(t *testing.T) {
			res := NewMaxIntHeap(test.In)
			if !reflect.DeepEqual(test.Expectation, res) {
				t.Error("Invalid result", test.Expectation, res)
			}
		})
	}

	addCases := []struct {
		Label       string
		Heap        MaxIntHeap
		In          int
		Expectation MaxIntHeap
	}{
		{
			"WithANewMax",
			MaxIntHeap{593, 49, 10, 9, 3},
			770,
			MaxIntHeap{770, 593, 49, 9, 3, 10},
		},
	}

	for _, test := range addCases {
		t.Run(test.Label, func(t *testing.T) {
			res := test.Heap.Add(test.In)

			if !reflect.DeepEqual(res, test.Expectation) {
				t.Error("Invalid result", test.Expectation, res)
			}
		})
	}

	popCases := []struct {
		Label        string
		Heap         MaxIntHeap
		ExpectedPop  int
		ExpectedHeap MaxIntHeap
	}{
		{
			"NominalCase",
			MaxIntHeap{770, 593, 49, 9, 3, 10},
			770,
			MaxIntHeap{593, 10, 49, 9, 3},
		},
	}

	for _, test := range popCases {
		t.Run(test.Label, func(t *testing.T) {
			popRes, popHeap := test.Heap.Pop()

			if popRes != test.ExpectedPop {
				t.Error("Invalid popped value", test.ExpectedPop, popRes)
			}

			if !reflect.DeepEqual(popHeap, test.ExpectedHeap) {
				t.Error("Invalid heap", test.ExpectedHeap, popHeap)
			}
		})
	}

}
