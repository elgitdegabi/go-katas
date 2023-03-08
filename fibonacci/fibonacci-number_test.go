package fibonacci

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test cases for FibonacciNumber
var fibonacciNumberBasicScenarios = []struct {
	value    int
	expected big.Int
}{
	{0, *big.NewInt(0)},
	{1, *big.NewInt(1)},
	{2, *big.NewInt(1)},
	{5, *big.NewInt(5)},
	{7, *big.NewInt(13)},
	{9, *big.NewInt(34)},
}

var fibonacciNumberComplexScenarios = []struct {
	value    int
	expected big.Int
}{
	{45, *big.NewInt(1134903170)},
	{85, *big.NewInt(259695496911122585)},
}

func Test_FibonacciNumber(t *testing.T) {
	for _, test := range fibonacciNumberBasicScenarios {
		assert.Equal(t, test.expected, calculate(test.value))
	}
}

func Test_FibonacciNumberWithMemoization(t *testing.T) {
	for _, test := range append(fibonacciNumberBasicScenarios, fibonacciNumberComplexScenarios...) {
		assert.Equal(t, test.expected, calculateWithMemoization(test.value, make(map[int]big.Int)))
	}
}
