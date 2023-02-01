package evenorodd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for EvenOrOdd
var scenarios = []struct {
	value    int
	expected string
	message  string
}{
	{0, "Even", "is not even"},
	{-1, "Odd", "is not odd"},
	{1, "Odd", "is not odd"},
	{-2, "Even", "is not even"},
	{2, "Even", "is not even"},
}

func Test_EvenOrOdd(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, evenOrOdd(test.value), test.message)
	}
}
