package fibonacci

import "math/big"

/**
 * Based on Interview challenge:
 * Calculate Nth Fibonacci number and optimize your solution to support "big numbers"
 */

const MAX_SUPPORTED_VALUE = 99

// Returns fibonacci Nth number
func calculate(k int) big.Int {
	if k < 1 || k > MAX_SUPPORTED_VALUE {
		return *big.NewInt(0)
	}
	if k <= 2 {
		return *big.NewInt(1)
	}

	previous := calculate(k - 1)
	beforePrevious := calculate(k - 2)

	return *big.NewInt(0).Add(&previous, &beforePrevious)
}

// Returns fibonacci Nth number using memoization
func calculateWithMemoization(k int, customCache map[int]big.Int) big.Int {
	if k < 1 || k > MAX_SUPPORTED_VALUE {
		return *big.NewInt(0)
	}
	if value, ok := customCache[k]; ok {
		return value
	}
	if k <= 2 {
		baseCase := *big.NewInt(1)
		customCache[k] = baseCase
		return baseCase
	}

	previous := calculateWithMemoization(k-1, customCache)
	beforePrevious := calculateWithMemoization(k-2, customCache)
	current := *big.NewInt(0).Add(&previous, &beforePrevious)

	customCache[k] = current

	return current
}
