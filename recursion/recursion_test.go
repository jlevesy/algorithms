package recursion

import (
	"reflect"
	"testing"
)

func TestCountPathTo(t *testing.T) {
	cases := []struct {
		Label       string
		Target      *GridPoint
		Expectation int
	}{
		{"NominalCase", &GridPoint{2, 2}, 6},
		{"InvalidEnd", &GridPoint{-2, 2}, 0},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			res := CountPathsTo(&GridPoint{0, 0}, test.Target)

			if res != test.Expectation {
				t.Error("Invalid result for", test.Target, test.Expectation, res)
			}
		})
	}
}

func TestFindPathTo(t *testing.T) {
	cases := []struct {
		Label         string
		Target        *GridPoint
		BlockedPoints map[GridPoint]struct{}
		Expectation   []GridPoint
	}{
		{
			"WithoutBlockingPoint",
			&GridPoint{2, 2},
			map[GridPoint]struct{}{},
			[]GridPoint{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			"WithBlockingPointAndASolution",
			&GridPoint{2, 2},
			map[GridPoint]struct{}{
				{2, 0}: {},
			},
			[]GridPoint{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}},
		},
		{
			"WithBlockingPointAndNoSolution",
			&GridPoint{2, 2},
			map[GridPoint]struct{}{
				{0, 1}: {},
				{1, 0}: {},
			},
			[]GridPoint{},
		},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			res := FindPathTo(&GridPoint{0, 0}, test.Target, test.BlockedPoints)

			if !reflect.DeepEqual(res, test.Expectation) {
				t.Error("Invalid result for", test.Target, test.Expectation, res)
			}
		})
	}
}
