package twooldestages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for two oldest ages
var scenarios = []struct {
	value    []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{}},
	{[]int{0, 0}, []int{0, 0}},
	{[]int{5, 10}, []int{5, 10}},
	{[]int{3, 1, 5}, []int{3, 5}},
	{[]int{1, 3, 10, 0}, []int{3, 10}},
	{[]int{1, 2, 10, 8}, []int{8, 10}},
	{[]int{1, 5, 87, 45, 8, 8}, []int{45, 87}},
	{[]int{2, 5, 87, 45, 8, 8}, []int{45, 87}},
	{[]int{6, 5, 83, 5, 3, 18}, []int{18, 83}},
	{[]int{6, 5, 83, 5, 3, 18}, []int{18, 83}},
}

func Test_TwoOldestAges(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, two_oldest_ages(test.value))
	}
}
