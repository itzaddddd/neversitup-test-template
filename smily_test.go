package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSime(t *testing.T) {
	testcases := []struct {
		Name   string
		Input  []string
		Expect int
	}{
		{
			Name:   "case [':)', ';(', ';}', ':-D']",
			Input:  []string{":)", ";(", ";}", ":-D"},
			Expect: 2,
		},
		{
			Name:   "case [';D', ':-(', ':-)', ';~)']",
			Input:  []string{";D", ":-(", ":-)", ";~)"},
			Expect: 3,
		},
		{
			Name:   "case [';]', ':[', ';*', ':$', ';-D']",
			Input:  []string{";]", ":[", ";*", ":$", ";-D"},
			Expect: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			output := countSmileys(tc.Input)

			assert.Equal(t, tc.Expect, output)

		})
	}
}
