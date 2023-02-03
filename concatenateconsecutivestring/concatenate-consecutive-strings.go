package concatenateconsecutivestring

/**
 * Based on Code-wars kata:
 * https://www.codewars.com/kata/56a5d994ac971f1ac500003e/go
 */

// concatenated k consecutive strings of length k or grater
func consecutiveString(strarr []string, k int) string {
	result := ""

	if len(strarr) > 0 && k > 0 && k < len(strarr) {
		for i := 0; i < len(strarr)-k+1; i++ {
			current := ""
			for j := i; j < i+k; j++ {
				current += strarr[j]
			}
			if len(current) > len(result) {
				result = current
			}
		}
	}

	return result
}
