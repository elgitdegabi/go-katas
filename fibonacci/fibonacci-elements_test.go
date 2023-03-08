package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for FibonacciElements
var fibonacciElementsScenarios = []struct {
	value    int
	expected []int
}{
	{0, []int{}},
	{1, []int{1}},
	{2, []int{1, 1}},
	{85, []int{}},
	{5, []int{1, 1, 2, 3, 5}},
	{7, []int{1, 1, 2, 3, 5, 8, 13}},
	{9, []int{1, 1, 2, 3, 5, 8, 13, 21, 34}},
}

func Test_FibonacciElements(t *testing.T) {
	for _, test := range fibonacciElementsScenarios {
		assert.Equal(t, test.expected, getElements(test.value))
	}
}
