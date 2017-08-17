package ds

// Node is a node of a linked list
type Node struct {
	value int32
	next  *Node
}

// Insert inserts a node in a given list at a given index
func Insert(head *Node, index int, node *Node) *Node {
	if head == nil {
		return node
	}

	if index == 0 {
		node.next = head
		return node
	}

	prev := At(head, index-1)

	if prev == nil {
		return nil
	}

	next := prev.next

	if next == nil {
		return nil
	}

	prev.next = node
	node.next = next

	return head
}

// At retrieves a node at a given index
// Time complexity: O(n)
// Space complexity: O(1)
func At(head *Node, index int) *Node {
	if head == nil {
		return nil
	}

	i := 0
	node := head

	for i < index && node != nil {
		node = node.next
		i++
	}

	return node
}

// Length returns the length of a list
// Time complexity: O(n)
// Space complexity: O(1)
func Length(head *Node) int {
	node := head
	length := 0

	for node != nil {
		length++
		node = node.next
	}

	return length
}

// Delete deletes a node from a linked list at a given index
// Time complexity: O(n)
// Space complexity: O(1)
func Delete(head *Node, index int) *Node {
	if head == nil {
		return nil
	}

	if index == 0 {
		return head.next
	}

	prev := At(head, index-1)

	if prev == nil {
		return head
	}

	node := prev.next

	if node == nil {
		return head
	}

	prev.next = node.next

	return head
}

// RemoveDuplicates removes duplicates from an unsorted linked list
// Questions :
// - Structure of a node : see above
// - Maximum value of a node : 0 <= x < 1000
// - Is there any loops ? Nope
// Idea n°1
// Scan for items, group by values, remove all instances with > 1 value
// - Time complexity O(n) worst case
// - Space complexity O(k) k different values
// Idea n°2
// For each item in the loop compare with all other items exept itself
// - Time complexity O(n²)
// - Space complexity O(1)
func RemoveDuplicates(list *Node) *Node {
	if list == nil {
		return nil
	}

	valueMap := map[int32]struct{}{
		list.value: {},
	}

	prev := list
	node := list.next

	for node != nil {
		if _, ok := valueMap[node.value]; ok {
			prev.next = node.next
			node = prev
		}

		valueMap[node.value] = struct{}{}

		prev = node
		node = node.next
	}

	return list
}

// Back implements an algorithm to find the kth to last element of
// a singly linked list
// - Loop ? Nope
// - Do we have length ? haha, nope
// - Do we have maximum length of the list ? nope
// - Can we alter state of a node ? nope
// Idea 1: Reverse the list then go the kth item
// - Time complexity O(n)
// - Space complexity O(n)
// Idea 2: Get the length of the list, go to the n-1-kth item
// - Time complexity O(2n) (worst case)
// - Space complexity O(1)
func Back(head *Node, k int) *Node {
	length := Length(head)
	return At(head, length-k-1)
}

// BackRecursive deals with the same problem but does it recursively
// - Time O(n)
// - Space O(n)
func BackRecursive(head *Node, k int) *Node {
	node, _ := backRecursion(head, k)
	return node
}

func backRecursion(head *Node, k int) (*Node, int) {
	if head == nil {
		return nil, 0
	}

	node, count := backRecursion(head.next, k)

	if node != nil {
		return node, 0
	}

	if count == k {
		return head, 0
	}

	count++

	return nil, count
}

// DeleteNode implements an algorithm to delete a node in the middle of a singly linked list
// given only access to that node
// This problem cannot be dealt with if we try to delete the last node
// Idea 1: Recurse until the end of the list,
// then copy value from current to previous
// Idea 2: Why do we need a recursion
// Time complexity: O(n) (n is length from node to end)
// Space complexity: O(n) (recursion, d'uh)
func DeleteNode(node *Node) {
	deleteNodeRecursion(node.next, node)
}

func deleteNodeRecursion(current, previous *Node) bool {
	if current == nil {
		return true
	}

	currentValue := current.value

	last := deleteNodeRecursion(current.next, current)

	previous.value = currentValue

	if last {
		previous.next = nil
	}

	return false
}

// DeleteNodeBetter is a better implementation of an algorithm
// to delete a node in the middle of a singly linked list
// given only access to that node
// Time complexity: O(n) (n is length from node to end)
// Space complexity: O(1)
func DeleteNodeBetter(node *Node) {
	for node != nil {
		if node.next != nil {
			node.value = node.next.value

			if node.next.next == nil {
				node.next = nil
			}
		}

		node = node.next
	}
}

// PartitionAroundValue partitions a linked list around a specific value
// such than all nodes < value are before and all nodes > value are after
// Questions asked:
// - Maximum value ? Infinity
// - Is there any loops ? Nope
// - Multiple node with the same value ? yes
// - Is the value mandatorily present in the list ? Nope
// - Is the sections supposed to be sorted ? nope
// - Am I authorized to mutate the old list ? Yes
// Idea 1: build three lists
// - 1 with nodes lower
// - 1 with nodes equals
// - 1 with nodel bigger
// Then concat and return the new head
func PartitionAroundValue(head *Node, value int32) *Node {
	var current, lower, lastLower, equal, lastEqual, bigger, lastBigger, list, lastList *Node
	current = head

	for current != nil {
		if current.value < value {
			lower, lastLower = appendList(lower, lastLower, current, current)
		} else if current.value == value {
			equal, lastEqual = appendList(equal, lastEqual, current, current)
		} else {
			bigger, lastBigger = appendList(bigger, lastBigger, current, current)
		}

		current = current.next
	}

	list, lastList = lower, lastLower

	if equal != nil {
		list, lastList = appendList(list, lastList, equal, lastEqual)
	}

	if bigger != nil {
		list, lastList = appendList(list, lastList, bigger, lastBigger)
	}

	lastList.next = nil

	return list
}

func appendList(head, last, appendHead, appendLast *Node) (*Node, *Node) {
	if head == nil {
		head, last = appendHead, appendLast
	} else {
		last.next = appendHead
		last = appendLast
	}

	return head, last
}
