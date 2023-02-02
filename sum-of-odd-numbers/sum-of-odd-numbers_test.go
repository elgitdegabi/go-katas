package sumofoddnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for EvenOrOdd
var scenarios = []struct {
	value    int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 8},
	{3, 27},
	{4, 64},
	{5, 125},
}

func Test_SumOddNumbersTriangle(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, sumOddNumbersTriangle(test.value))
	}
}
