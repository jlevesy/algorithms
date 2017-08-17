package ds

import (
	"testing"
)

func equals(l1, l2 *Node) bool {
	n1, n2 := l1, l2
	for n1 != nil && n2 != nil {
		if n2 == nil || n1.value != n2.value {
			return false
		}

		n1 = n1.next
		n2 = n2.next
	}

	if n1 != n2 {
		return false
	}

	return true
}

func TestLinkedList(t *testing.T) {
	insertCases := []struct {
		Label       string
		List        *Node
		Index       int
		Node        *Node
		Expectation *Node
	}{
		{
			"NominalCase",
			&Node{12, &Node{14, &Node{5, nil}}},
			1,
			&Node{5, nil},
			&Node{12, &Node{5, &Node{14, &Node{5, nil}}}},
		},
	}

	for _, test := range insertCases {
		t.Run(test.Label, func(t *testing.T) {
			res := Insert(test.List, test.Index, test.Node)

			if !equals(res, test.Expectation) {
				t.Error("Unexpected result for insert")
			}
		})
	}

	deleteCases := []struct {
		Label       string
		List        *Node
		Index       int
		Expectation *Node
		DelNode     *Node
	}{
		{
			"NominalCase",
			&Node{12, &Node{14, &Node{5, nil}}},
			1,
			&Node{12, &Node{5, nil}},
			&Node{14, nil},
		},
	}

	for _, test := range deleteCases {
		t.Run(test.Label, func(t *testing.T) {
			list := Delete(test.List, test.Index)

			if !equals(list, test.Expectation) {
				t.Error("Unexpected result for insert")
			}
		})
	}

	removeDupCases := []struct {
		Label       string
		List        *Node
		Expectation *Node
	}{
		{
			"WithMultipleDuplicates",
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
			&Node{10, &Node{5, &Node{12, &Node{32, nil}}}},
		},
		{
			"WithDuplicatesOfFirstItem",
			&Node{10, &Node{10, &Node{10, &Node{10, &Node{10, nil}}}}},
			&Node{10, nil},
		},
	}

	for _, test := range removeDupCases {
		t.Run(test.Label, func(t *testing.T) {
			list := RemoveDuplicates(test.List)

			if !equals(list, test.Expectation) {
				t.Error("Unexpected result for remove duplicates")
			}
		})
	}

	backCases := []struct {
		Label       string
		List        *Node
		K           int
		Expectation *Node
	}{
		{
			"NominalCase",
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
			1,
			&Node{5, nil},
		},
	}

	for _, test := range backCases {
		t.Run(test.Label, func(t *testing.T) {
			res := Back(test.List, test.K)
			res2 := BackRecursive(test.List, test.K)

			if res.value != test.Expectation.value {
				t.Error("Unexpected result for back")
			}

			if res2.value != test.Expectation.value {
				t.Error("Unexpected result for back2")
			}
		})
	}

	secondPart := &Node{10, &Node{5, nil}}
	list := &Node{10, &Node{5, &Node{12, &Node{5, &Node{32, secondPart}}}}}
	expectation := &Node{10, &Node{5, &Node{12, &Node{5, &Node{32, &Node{5, nil}}}}}}

	DeleteNodeBetter(secondPart)

	if !equals(list, expectation) {
		t.Error("DeleteNode failed")
	}

	partitionCase := []struct {
		Label       string
		List        *Node
		Value       int32
		Expectation *Node
	}{
		{
			"NominalCase",
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
			12,
			&Node{10, &Node{5, &Node{5, &Node{12, &Node{32, nil}}}}},
		},
		{
			"WithMaxValue",
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
			32,
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
		},
		{
			"WithNotPresentValue",
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
			15,
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
		},
		{
			"WithMultipleEqualsValue",
			&Node{10, &Node{5, &Node{12, &Node{5, &Node{32, nil}}}}},
			5,
			&Node{5, &Node{5, &Node{10, &Node{12, &Node{32, nil}}}}},
		},
	}

	for _, test := range partitionCase {
		t.Run(test.Label, func(t *testing.T) {
			res := PartitionAroundValue(test.List, test.Value)

			if !equals(res, test.Expectation) {
				t.Error("Unexpected result for PartitionAroundValue")
			}
		})
	}

}
