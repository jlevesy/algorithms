package ds

// MaxIntHeap is an implementation of an integer max heap
type MaxIntHeap []int

// NewMaxIntHeap returns a MaxIntHeap given a random int slice
func NewMaxIntHeap(in []int) MaxIntHeap {
	res := make(MaxIntHeap, 0, len(in))

	for _, item := range in {
		res = res.Add(item)
	}

	return res
}

// Add add an item to the heap
func (m MaxIntHeap) Add(item int) MaxIntHeap {
	m = append(m, item)
	m.heapifyUp(len(m) - 1)
	return m
}

// Pop the first item then maintain the heap
func (m MaxIntHeap) Pop() (int, MaxIntHeap) {
	item := m[0]
	m[0] = m[len(m)-1]
	m.heapifyDown(0)
	return item, m[0 : len(m)-1]
}

func (m MaxIntHeap) heapifyDown(i int) {
	l := left(i)
	r := right(i)

	for m[i] < m[l] || m[i] < m[r] {
		if m[i] < m[l] {
			m[i], m[l] = m[l], m[i]
			i = l
		}

		if m[i] < m[r] {
			m[i], m[r] = m[r], m[i]
			i = r
		}

		l = left(i)
		r = right(i)
	}
}

func (m MaxIntHeap) heapifyUp(i int) {
	p := parent(i)
	for m[i] > m[p] {
		m[p], m[i] = m[i], m[p]
		i = p
		p = parent(i)
	}
}

func left(index int) int {
	return (index + 1) / 2
}

func right(index int) int {
	return (index + 2) / 2
}

func parent(index int) int {
	return index / 2
}
