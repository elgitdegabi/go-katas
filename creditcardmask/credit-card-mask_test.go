package creditcardmask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for credit card mask
var scenarios = []struct {
	value    string
	expected string
}{
	{"", ""},
	{"abc", "abc"},
	{"abcd", "abcd"},
	{"4556364607935616", "############5616"},
	{"64607935616", "#######5616"},
	{"Skippy", "##ippy"},
	{"Nananananananananananananananana Batman!", "####################################man!"},
}

func Test_Maskify(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, maskify(test.value))
	}
}
