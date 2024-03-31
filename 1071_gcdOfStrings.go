package main

// gcdOfStrings returns the greatest common divisor of two strings.
// It returns an empty string if the strings have no common patterns.
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 { // if concatenation of str1 and str2 is not equal to str2+str1, it means that we cannot find gcd (i.e. they do not have common patterns)
		return ""
	}

	return str1[:gcd(len(str1), len(str2))] // find gcd and return the line from the beginning of the line to gcd
}

// GCD returns the greatest common divisor of two integers.
func gcd(a int, b int) int { // Euclidean algorithm for finding gcd
	for a != 0 && b != 0 { // if a division in b without remain, it means that smallest number is gcd, else if there is remain, the bigger number replace in remain
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return a + b
}
