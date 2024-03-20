package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOddNumberShouldBeCorrect(t *testing.T) {
	// TODO : start your code here
	assert := assert.New(t)

	cases := []struct {
		input    []int
		expected int
	}{
		{input: []int{7}, expected: 7},
		{input: []int{0}, expected: 0},
		{input: []int{1, 1, 2}, expected: 2},
		{input: []int{0, 1, 0, 1, 0}, expected: 0},
		{input: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, expected: 4},
	}

	for _, c := range cases {
		actual := FindOddNumber(c.input)
		assert.Equal(c.expected, actual)
	}

}
