package twooldestages

import "sort"

/**
 * Based on Code-wars kata:
 * https://www.codewars.com/kata/511f11d355fe575d2c000001/go
 */

// Gets two oldest ages from given array of ages
func two_oldest_ages(ages []int) []int {
	if len(ages) < 2 {
		return []int{}
	}

	sort.Ints(ages)
	return []int{ages[len(ages)-2], ages[len(ages)-1]}
}
