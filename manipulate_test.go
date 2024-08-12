package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	testcases := []struct {
		Name   string
		Input  string
		Expect []string
	}{
		{
			Name:   "case 'a'",
			Input:  "a",
			Expect: []string{"a"},
		},
		{
			Name:   "case 'ab'",
			Input:  "ab",
			Expect: []string{"ab", "ba"},
		},
		{
			Name:   "case 'abc'",
			Input:  "abc",
			Expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			Name:   "case 'aabb'",
			Input:  "aabb",
			Expect: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			output := permute(tc.Input)

			assert.Equal(t, tc.Expect, output)
		})
	}
}
