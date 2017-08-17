package ds

import (
	"testing"
)

func TestHasPathTo(t *testing.T) {
	nodeA := &GraphNode{ID: 'A'}
	nodeB := &GraphNode{ID: 'B'}
	nodeC := &GraphNode{ID: 'C'}
	nodeD := &GraphNode{ID: 'D'}
	nodeE := &GraphNode{ID: 'E'}
	nodeF := &GraphNode{ID: 'F'}

	nodeA.AddLinks(nodeB, nodeD)
	nodeB.AddLinks(nodeC, nodeE)
	nodeC.AddLinks(nodeA, nodeF)
	nodeD.AddLinks(nodeE)

	cases := []struct {
		Label       string
		StatingNode *GraphNode
		Target      rune
		Expectation bool
	}{
		{"WithPath", nodeA, 'E', true},
		{"WithNoPath", nodeD, 'C', false},
		{"WithSameRoot", nodeA, 'A', true},
		{"WithNotExistingNode", nodeA, 'Z', false},
		{"WithNilEntry", nil, 'A', false},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			resDFS := HasPathToDFS(test.StatingNode, test.Target)
			resBFS := HasPathToBFS(test.StatingNode, test.Target)

			if resDFS != test.Expectation {
				t.Error("Invalid result for DFS in case", test.Label)
			}

			if resBFS != test.Expectation {
				t.Error("Invalid result for BFS in case", test.Label)
			}
		})
	}
}
