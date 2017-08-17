package ds

import (
	"testing"
)

func equals(l1, l2 *LinkedListNode) bool {
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
		List        *LinkedListNode
		Index       int
		Node        *LinkedListNode
		Expectation *LinkedListNode
	}{
		{
			"NominalCase",
			&LinkedListNode{12, &LinkedListNode{14, &LinkedListNode{5, nil}}},
			1,
			&LinkedListNode{5, nil},
			&LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{14, &LinkedListNode{5, nil}}}},
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
		List        *LinkedListNode
		Index       int
		Expectation *LinkedListNode
		DelNode     *LinkedListNode
	}{
		{
			"NominalCase",
			&LinkedListNode{12, &LinkedListNode{14, &LinkedListNode{5, nil}}},
			1,
			&LinkedListNode{12, &LinkedListNode{5, nil}},
			&LinkedListNode{14, nil},
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
		List        *LinkedListNode
		Expectation *LinkedListNode
	}{
		{
			"WithMultipleDuplicates",
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{32, nil}}}},
		},
		{
			"WithDuplicatesOfFirstItem",
			&LinkedListNode{10, &LinkedListNode{10, &LinkedListNode{10, &LinkedListNode{10, &LinkedListNode{10, nil}}}}},
			&LinkedListNode{10, nil},
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
		List        *LinkedListNode
		K           int
		Expectation *LinkedListNode
	}{
		{
			"NominalCase",
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
			1,
			&LinkedListNode{5, nil},
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

	secondPart := &LinkedListNode{10, &LinkedListNode{5, nil}}
	list := &LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, secondPart}}}}}
	expectation := &LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, &LinkedListNode{5, nil}}}}}}

	DeleteNodeBetter(secondPart)

	if !equals(list, expectation) {
		t.Error("DeleteNode failed")
	}

	partitionCase := []struct {
		Label       string
		List        *LinkedListNode
		Value       int32
		Expectation *LinkedListNode
	}{
		{
			"NominalCase",
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
			12,
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{32, nil}}}}},
		},
		{
			"WithMaxValue",
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
			32,
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
		},
		{
			"WithNotPresentValue",
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
			15,
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
		},
		{
			"WithMultipleEqualsValue",
			&LinkedListNode{10, &LinkedListNode{5, &LinkedListNode{12, &LinkedListNode{5, &LinkedListNode{32, nil}}}}},
			5,
			&LinkedListNode{5, &LinkedListNode{5, &LinkedListNode{10, &LinkedListNode{12, &LinkedListNode{32, nil}}}}},
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
