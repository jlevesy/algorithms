package ds

// GraphNode is a node of a GraphNode
type GraphNode struct {
	ID    rune
	links []*GraphNode
}

// AddLinks adds a link to a node
func (g *GraphNode) AddLinks(node ...*GraphNode) {
	g.links = append(g.links, node...)
}

// HasPathTo
// Write a method which tells if there is a path from given
// root node to the node with given id
// Questions:
// - Is there any loops in the graph ? yes
// - Is there a maximum node count, no
// - Is there always a path ? no

// HasPathToDFS implements a solution using depth first search
// Idea:
// Check if the id is given node, if yes return true
// Maintain a set of visited ids, in order to detect loops
// If not itself for each link, recurse on the current child  if not already visited
// until found true
// Time complexity: O(n) worst case (n is the number of vertices)
// Space complexity: O(1)
func HasPathToDFS(root *GraphNode, id rune) bool {
	visitedSet := map[rune]struct{}{}
	return hasPathDFSRecursion(root, id, visitedSet)
}

func hasPathDFSRecursion(node *GraphNode, id rune, visitedSet map[rune]struct{}) bool {
	if node == nil {
		return false
	}

	visitedSet[node.ID] = struct{}{}

	if node.ID == id {
		return true
	}

	for _, child := range node.links {
		if _, ok := visitedSet[child.ID]; ok {
			continue
		}

		found := hasPathDFSRecursion(child, id, visitedSet)

		if found {
			return found
		}
	}

	return false
}

// HasPathToBFS implements a solution using breadth first search
// Idea:
// Create a queue, push back first node into the queue
// While the queue is not empty, check if popped node is the target,
// if not, push back all non already visited node childs to the queue
func HasPathToBFS(root *GraphNode, id rune) bool {
	if root == nil {
		return false
	}

	queue := graphNodeQueue{}
	visited := map[rune]struct{}{}

	queue.Push(root)

	for !queue.Empty() {
		node, ok := queue.Pop()

		if !ok {
			return false
		}

		if _, ok := visited[node.ID]; ok {
			continue
		}

		visited[node.ID] = struct{}{}

		if node.ID == id {
			return true
		}

		queue.Push(node.links...)
	}

	return false
}

type graphNodeQueue []*GraphNode

func (g *graphNodeQueue) Push(nodes ...*GraphNode) {
	*g = append(*g, nodes...)
}

func (g *graphNodeQueue) Empty() bool {
	return len(*g) == 0
}

func (g *graphNodeQueue) Pop() (*GraphNode, bool) {
	if g.Empty() {
		return nil, false
	}

	node, newG := (*g)[0], (*g)[1:]
	*g = newG

	return node, true
}
