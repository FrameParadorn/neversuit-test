package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManipuateShouldBeReturnOneElement(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"a"}, expected: []string{"a"}},
		{input: []string{"b"}, expected: []string{"b"}},
		{input: []string{"c"}, expected: []string{"c"}},
	}

	for _, c := range cases {
		actual := Manipulate(c.input)
		assert.Equal(c.expected, actual)
	}
}

func TestManipuateShouldBeReturnShuffingElementCorrect(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"a", "b"}, expected: []string{"ab", "ba"}},
		{input: []string{"a", "b", "c"}, expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{input: []string{"a", "a", "b", "b"}, expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, c := range cases {
		actual := Manipulate(c.input)
		fmt.Println(actual)
		assert.ElementsMatch(c.expected, actual)
	}
}
