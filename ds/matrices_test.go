package ds

import (
	"reflect"
	"testing"
)

func TestRotateMatrixRight(t *testing.T) {
	cases := []struct {
		Label       string
		In          [][]int32
		Expectation [][]int32
	}{
		{
			"NominalCase",
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int32{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			"With1x1Matrix",
			[][]int32{
				{1},
			},
			[][]int32{
				{1},
			},
		},
		{
			"WithEmptyMAtrix",
			[][]int32{},
			[][]int32{},
		},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			RotateMatrixRight(test.In)
			if !reflect.DeepEqual(test.In, test.Expectation) {
				t.Error("Invalid result for", test.In, test.Expectation)
			}
		})
	}
}

func TestZeroRowAndColumn(t *testing.T) {
	cases := []struct {
		Label       string
		In          [][]int32
		Expectation [][]int32
	}{
		{
			"NominalCase",
			[][]int32{
				{1, 2, 3, 59},
				{4, 5, 6, 298},
				{0, 8, 9, 0},
			},
			[][]int32{
				{0, 2, 3, 0},
				{0, 5, 6, 0},
				{0, 0, 0, 0},
			},
		},
		{
			"With1x1Matrix",
			[][]int32{
				{1},
			},
			[][]int32{
				{1},
			},
		},
		{
			"WithEmptyMAtrix",
			[][]int32{},
			[][]int32{},
		},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			ZeroRowAndColumn(test.In)
			if !reflect.DeepEqual(test.In, test.Expectation) {
				t.Error("Invalid result for", test.In, test.Expectation)
			}
		})
	}
}
