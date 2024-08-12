package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCountOdd(t *testing.T) {
	testcases := []struct {
		Name   string
		Input  []int
		Expect int
	}{
		{
			Name:   "case [7]",
			Input:  []int{7},
			Expect: 7,
		},
		{
			Name:   "case [0]",
			Input:  []int{0},
			Expect: 0,
		},
		{
			Name:   "case [1,1,2]",
			Input:  []int{1, 1, 2},
			Expect: 2,
		},
		{
			Name:   "case [0,1,0,1,0]",
			Input:  []int{0, 1, 0, 1, 0},
			Expect: 0,
		},
		{
			Name:   "case [1,2,2,3,3,3,4,3,3,3,2,2,1]",
			Input:  []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			Expect: 4,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			output := findCountOdd(tc.Input)

			assert.Equal(t, tc.Expect, output)
		})
	}
}
