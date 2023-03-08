package localminmax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for LocalMinMax
var scenarios = []struct {
	value    []int
	expected [][]int
}{
	{nil, [][]int{{0}, {0}}},
	{[]int{}, [][]int{{0}, {0}}},
	{[]int{1}, [][]int{{1}, {1}}},
	{[]int{1, 1}, [][]int{{1}, {1}}},
	{[]int{1, 2, 3, 4}, [][]int{{1}, {4}}},
	{[]int{4, 3, 2, 1}, [][]int{{1}, {4}}},
}

func Test_LocalMinMax(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, getLocalMinMax(test.value))
		assert.Equal(t, test.expected, getLocalMinMaxAlternative(test.value))
	}
}
