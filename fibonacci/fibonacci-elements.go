package fibonacci

/**
 * Based on Geeks for geeks exercise:
 * https://practice.geeksforgeeks.org/problems/print-first-n-fibonacci-numbers1002/1
 */

// Returns first K fibonacci elements
func getElements(k int) []int {
	if k < 1 || k > 84 {
		return []int{}
	}

	result := make([]int, k)

	for i := 0; i < k; i++ {
		if i < 2 {
			result[i] = 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}

	return result
}
