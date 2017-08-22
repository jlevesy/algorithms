package ds

import (
	"errors"
)

// Describe how you could use a single array to implement K stacks
// Questions:
// - Do we have a max stack size ? Nope
// - Do we have a max array size ? Let's say N items
// - Could the array grow ? Yes (and it is not an array then)
// - What are supported operations: Push(stackID) and Pop(stackID)
// - Am I authorized to keep indexes in stack ? indeed, using a slice ? yes
// Idea n°1:
// If we have a maximum stack size we could partition the undelying array
// in K segment of length N/K. Fairly straightforward however if the undelyinng
// array grows, it can create problems.
// Idea n°2:
// Instead, what we might want to do is to use 1 stack item each K items, this way the
// undelying array could grow, because we're not making any expectation on its size
// However it implies an equivalent usage of the three stacks. Or It would be vastly
// inneficient.
// Now we might want to have a look about how Push and Pop are supposed to work.
// - Push:
// Get stack ID last item index, increment +1 and  write at ID+stackID
// Complexity O(1), worst case O(N) if there is an array resizing
// - Pop:
// Get Stack ID last item, access the item, drop the reference to item, decrement stack counteter

// MultiStack is an attempt to implement that
type MultiStack struct {
	storage []uint32
	indexes []int
}

// NewMultiStack constructs a mutlistack
func NewMultiStack(stacks, size int) *MultiStack {
	return &MultiStack{
		storage: make([]uint32, size),
		indexes: make([]int, stacks),
	}
}

// Empty returns true if the given stack doesn't exist
// or exists and in empty. else false
func (m *MultiStack) Empty(stackID int) bool {
	if stackID > len(m.indexes) {
		return true
	}

	return m.indexes[stackID] == 0
}

// Pop pops and item from the stack with id stackID
func (m *MultiStack) Pop(stackID int) (uint32, error) {
	if stackID > len(m.indexes) {
		return 0, errors.New("Stack does not exist")
	}

	if m.indexes[stackID] == 0 {
		return 0, errors.New("Stack is empty")
	}

	m.indexes[stackID]--
	itemIndex := storageIndex(stackID, len(m.indexes), m.indexes[stackID])

	item := m.storage[itemIndex]
	m.storage[itemIndex] = 0

	return item, nil
}

// Push push an item onto the given stack
func (m *MultiStack) Push(stackID int, item uint32) error {
	if stackID > len(m.indexes) {
		return errors.New("Stack does not exist")
	}

	itemIndex := storageIndex(stackID, len(m.indexes), m.indexes[stackID])

	if itemIndex > len(m.storage) {
		m.doubleStorage()
	}

	m.storage[itemIndex] = item
	m.indexes[stackID]++

	return nil
}

func storageIndex(stackID, stacksCount, index int) int {
	return index*stacksCount + stackID
}

func (m *MultiStack) doubleStorage() {
	newBuf := make([]uint32, 2*len(m.storage))
	copy(newBuf, m.storage)
	m.storage = newBuf
}

// Hanoi problem
// - 3 towers of n disks
// - 1 tower with disks sorted in ascending order ( disk always sorted of a larger one)
// - Only one disk can be moved at a time
// - A disk slid off the top of one tower onto the next tower
// - A disk can be placed on top of a larger disk
// Write a program to move the disks from the first tower to the last using stacks.
