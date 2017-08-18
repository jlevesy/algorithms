package ds

import (
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	cases := []struct {
		Label       string
		Subject     *LinkedQueue
		Expectation bool
	}{
		{"WithEmptyQueue", &LinkedQueue{}, true},
		{"WithNonEmptyQueue", &LinkedQueue{&LinkedListNode{}, &LinkedListNode{}}, false},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			if test.Subject.Empty() != test.Expectation {
				t.Error("InvalidResult for T")
			}
		})
	}
}

func TestLinkedQueue(t *testing.T) {
	cases := []struct {
		Label string
		Nodes []*LinkedListNode
	}{
		{"NominalCase", []*LinkedListNode{{}, {}, {}}},
		{"EmptyNodes", []*LinkedListNode{}},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			subject := LinkedQueue{}
			subject.Enqueue(test.Nodes...)
			res := []*LinkedListNode{}
			for !subject.Empty() {
				node, ok := subject.Dequeue()
				if !ok {
					t.Error("Queue empty before expected")
					t.FailNow()
				}
				res = append(res, node)
			}

			// Comparing ptrs is good enough in that case
			if !reflect.DeepEqual(test.Nodes, res) {
				t.Error("Invalid result for Enqueue / Dequeue", test.Nodes, res)
			}
		})
	}
}
