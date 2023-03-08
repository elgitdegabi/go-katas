package localminmax

import "sort"

/**
 * Based on Google interview challenge
 */

// Get local min and max values from given int array
func getLocalMinMax(list []int) [][]int {
	result := [][]int{{0}, {0}}

	if len(list) > 0 {
		min := list[0]
		max := list[0]

		for _, value := range list {
			if value < min {
				min = value
			} else if value > max {
				max = value
			}
		}

		result[0][0] = min
		result[1][0] = max
	}

	return result
}

// Get local min and max values from given int array
func getLocalMinMaxAlternative(list []int) [][]int {
	result := [][]int{{0}, {0}}

	if len(list) > 0 {
		sort.Ints(list)
		result[0][0] = list[0]
		result[1][0] = list[len(list)-1]
	}

	return result
}
