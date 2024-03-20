package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmilyFaceShouldBeCorrect(t *testing.T) {
	// TODO : start your code here
	assert := assert.New(t)

	cases := []struct {
		input    []string
		expected int
	}{
		{input: []string{":)", ";(", ":-D"}, expected: 2},
		{input: []string{";D", ":-(", ":-)", ";~)"}, expected: 3},
		{input: []string{";]", ":[", ";*", ";-D"}, expected: 1},
	}

	for _, c := range cases {
		actual := CountSmilyFace(c.input)
		assert.Equal(c.expected, actual)
	}

}
