package splitstringintopair

/**
 * Based on Code-wars kata:
 * https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/go
 */

// Splits given string into pairs. If original string has odd length, it should be completed with _ character
func split(str string) []string {
	if len(str) < 1 {
		return []string{}
	}

	if len(str)%2 != 0 {
		str += "_"
	}

	result := make([]string, len(str)/2)

	for i := 0; i < len(result); i++ {
		result[i] = str[i*2 : i*2+2]
	}

	return result
}
