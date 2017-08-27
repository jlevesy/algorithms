package recursion

// Imagine a robot sitting on the upper left corner of an X by Y grid,
// the robot can move down only to the right, or down.
// How many possible paths are there for the robot to go from 0,0 to X,Y
// Questions:
// - Maximum size of the grid ? None
// Examples
// o,_,_
// _,_,_ =>  (X=2, Y=2) = 6 paths possible
// _,_,x
// - 0,0 => 1,0 => 2,0 => 2,1 => 2,2
// - 0,0 => 1,0 => 1,1 => 2,1 => 2,2
// - 0,0 => 1,0 => 1,1 => 1,2 => 2,2
// - 0,0 => 0,1 => 1,1 => 2,1 => 2,2
// - 0,0 => 0,1 => 1,1 => 1,2 => 2,2
// - 0,0 => 0,1 => 0,2 => 1,2 => 2,2

// GridPoint represents a point in a grid
type GridPoint struct {
	X int
	Y int
}

func (p *GridPoint) right() *GridPoint {
	return &GridPoint{p.X + 1, p.Y}
}

func (p *GridPoint) down() *GridPoint {
	return &GridPoint{p.X, p.Y + 1}
}

// CountPathsTo is an implementation of this problem
func CountPathsTo(start, target *GridPoint) int {
	if *start == *target {
		return 1
	}

	if start.X > target.X || start.Y > target.Y {
		return 0
	}

	return CountPathsTo(start.right(), target) +
		CountPathsTo(start.down(), target)
}

// Follow up to the previous problem, write an algorithm which finds a pathRight
// It is supposed to take into account various blockedPoints

// FindPathTo is a DFS algorithm variant retuning one path to the target point
// Time and space complexity are O(2^n) n is the number of steps necessary to find a path
func FindPathTo(start, target *GridPoint, blockedPoints map[GridPoint]struct{}) []GridPoint {
	if *start == *target {
		return []GridPoint{*target}
	}

	if start.X > target.X || start.Y > target.Y {
		return []GridPoint{}
	}

	if _, ok := blockedPoints[*start]; ok {
		return []GridPoint{}
	}

	pathRight := FindPathTo(start.right(), target, blockedPoints)

	if len(pathRight) != 0 {
		return append([]GridPoint{*start}, pathRight...)
	}

	pathDown := FindPathTo(start.down(), target, blockedPoints)

	if len(pathDown) != 0 {
		return append([]GridPoint{*start}, pathDown...)
	}

	return []GridPoint{}
}

// Implement an algorithm to print all valid (properly opened and closed) combinations
// of n pairs of parentheses.

// Questions asked:
// - ...? Problem seems pretty clear to me

// Idea
// Sounds like a recursion problem.
// We notice that to create N == 2, we insert a parenthesis pairs
// at each possible position for each previous solution.
// It can lead to duplicates, so we deduplicate at the end.

// Examples

// N == 3 => ["()()()", "(())()", "()(())", ((()))"]
// N == 2 => ["()()", "(())"]
// N == 1 => ["()"]
// N == 0 => []

// GenerateParenthesis is an implementation of that
func GenerateParenthesis(pairs int) []string {
	res := []string{}
	index := map[string]struct{}{}
	generated := generateParenthesisRecursion(pairs, 0, 0)

	for _, s := range generated {
		if _, ok := index[s]; ok {
			continue
		}

		index[s] = struct{}{}
		res = append(res, s)
	}

	return res
}

func insertAt(in, chunk string, position int) string {
	return in[:position] + chunk + in[position:]
}

func generateParenthesisRecursion(pairs, openCount, CloseCount int) []string {
	res := []string{}

	if pairs == 0 {
		return res
	}

	if pairs == 1 {
		return []string{"()"}
	}

	previousLevel := generateParenthesisRecursion(pairs-1, 0, 0)

	for _, s := range previousLevel {
		for index := range s {
			res = append(res, insertAt(s, "()", index))
		}
	}

	return res
}
