package concatenateconsecutivestring

import (
	"reflect"
)

/**
 * Based on Code-wars kata:
 * https://www.codewars.com/kata/56a5d994ac971f1ac500003e/go
 */

// concatenated k consecutive strings of length k or grater
func consecutiveString(array interface{}, k int) string {
	strarr := reflect.ValueOf(array)
	result := ""

	if strarr.Len() > 0 && k > 0 && k < strarr.Len() {
		for i := 0; i < strarr.Len()-k+1; i++ {
			current := ""
			for j := i; j < i+k; j++ {
				current += strarr.Index(j).String()
			}
			if len(current) > len(result) {
				result = current
			}
		}
	}

	return result
}
