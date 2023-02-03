package concatenateconsecutivestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for consecutiveString
var scenarios = []struct {
	array    []string
	k        int
	expected string
}{
	{[]string{"a", "b", "cd", "efg"}, 2, "cdefg"},
	{[]string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}, 2, "folingtrashy"},
	{[]string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"}, 2, "abigailtheta"},
	{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1, "oocccffuucccjjjkkkjyyyeehh"},
	{[]string{}, 3, ""},
	{[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2, "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck"},
	{[]string{"wlwsasphmxx", "owiaxujylentrklctozmymu", "wpgozvxxiu"}, 2, "wlwsasphmxxowiaxujylentrklctozmymu"},
	{[]string{"zone", "abigail", "theta", "form", "libe", "zas"}, -2, ""},
	{[]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}, 3, "ixoyx3452zzzzzzzzzzzz"},
	{[]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}, 15, ""},
	{[]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}, 0, ""},
}

func Test_ConsecutiveString(t *testing.T) {
	for _, test := range scenarios {
		assert.Equal(t, test.expected, consecutiveString(test.array, test.k))
	}
}
