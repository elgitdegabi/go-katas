package splitstringintopair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for split string into pair
var scenarios = []struct {
	value    string
	expected []string
}{
	{"", []string{}},
	{"a", []string{"a_"}},
	{"ab", []string{"ab"}},
	{"abc", []string{"ab", "c_"}},
	{"abcd", []string{"ab", "cd"}},
}

func Test_Split(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, split(test.value))
	}
}
