package creditcardmask

/**
 * Based on Code-wars kata:
 * https://www.codewars.com/kata/5412509bd436bd33920011bc/solutions/go
 */

// Generates a mask for given str except last 4 characters
func maskify(str string) string {
	if len(str) > 4 {
		mask := ""
		for i := 0; i < len(str)-4; i++ {
			mask += "#"
		}

		return mask + str[len(str)-4:]
	}

	return str
}
