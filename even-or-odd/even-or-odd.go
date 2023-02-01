package evenorodd

/**
 * Based on Code-wars kata:
 * https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/go
 */

//  Returns if given number is even or odd
func evenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
