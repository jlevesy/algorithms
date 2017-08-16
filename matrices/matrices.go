package matrices

// RotateMatrixRight rotates given matrix 90 degrees to the right
// Time O(n^2)
// Space O(1)
func RotateMatrixRight(m [][]int32) {
	for layer := 0; layer < len(m)/2; layer++ {
		first, last := layer, len(m[layer])-1-layer
		for i := first; i < last; i++ {
			offset := i - first
			top := m[first][i]
			// left => top
			m[first][i] = m[last-offset][first]

			// bottom => left
			m[last-offset][first] = m[last][last-offset]

			// right => bottom
			m[last][last-offset] = m[i][last]

			// top => right
			m[i][last] = top

		}
	}
}

// ZeroRowAndColumn zeroes row and column of a MxN matrix  if the value is 0
// - Do we expect to override zeroes ?  Nope
// - Should it be done inplace ? Not necessarily
// - Is it a square matrix ? Nope
// - What are my inputs ? The matrix
// Idea n°1: Allocate a matrix of zeroed MxN values,
// Scan for zeroes, store their coordinates in a set
// If new coordinates matches the x or y of any spotted zeroes, just don't write the value
// else write the value
// Idea n°2 keep the scanning session, then zero each spotted row and colomn
// Time O(M*N)
// Space O(s) s ==  count of spotted zeroes
func ZeroRowAndColumn(m [][]int32) {
	if len(m) == 0 || len(m[0]) == 0 {
		return
	}

	zeroesRows := map[int]struct{}{}
	zeroesColumns := map[int]struct{}{}

	for rowIndex, row := range m {
		for colIndex, item := range row {
			if item == 0 {
				zeroesRows[rowIndex] = struct{}{}
				zeroesColumns[colIndex] = struct{}{}
			}

		}
	}

	for rowIndex := range zeroesRows {
		for i := range m[rowIndex] {
			m[rowIndex][i] = 0
		}
	}

	for colIndex := range zeroesColumns {
		for _, row := range m {
			row[colIndex] = 0
		}
	}
}
