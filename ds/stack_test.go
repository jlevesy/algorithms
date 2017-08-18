package ds

import (
	"reflect"
	"testing"
)

func TestMultiStack(t *testing.T) {
	cases := []struct {
		Label       string
		StorageSize int
		Stacks      [][]uint32
		Expectation [][]uint32
	}{
		{
			"NominalCase",
			9,
			[][]uint32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]uint32{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
		},
		{
			"WithRealocation",
			5,
			[][]uint32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]uint32{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
		},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			subject := NewMultiStack(len(test.Stacks), test.StorageSize)
			for stackID, stack := range test.Stacks {
				for _, item := range stack {
					if err := subject.Push(stackID, item); err != nil {
						t.Error("Were'nt supposed to fail here")
						t.FailNow()
					}
				}
			}

			results := make([][]uint32, len(test.Stacks))

			for stackID := range results {
				for !subject.Empty(stackID) {
					item, err := subject.Pop(stackID)
					if err != nil {
						t.Error("Were'nt supposed to fail here")
						t.FailNow()
					}
					results[stackID] = append(results[stackID], item)
				}
			}

			for stackID, stack := range results {
				if !reflect.DeepEqual(stack, test.Expectation[stackID]) {
					t.Error("Invalid result  for stack", stackID, stack, test.Expectation[stackID])
				}
			}
		})
	}
}
