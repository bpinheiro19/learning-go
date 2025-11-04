package main

import (
	"fmt"
)

/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.

The algorithm for myAtoi(string s) is as follows:
	Whitespace: Ignore any leading whitespace (" ").
	Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
	Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
	Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
Return the integer as the final result.

Example 1:
Input: s = "42"
Output: 42

Example 2:
Input: s = " -042"
Output: -42

Example 3:
Input: s = "1337c0d3"
Output: 1337

Example 4:
Input: s = "0-1"
Output: 0

Example 5:
Input: s = "words and 987"
Output: 0


Constraints:
0 <= s.length <= 200
s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
*/

func main() {
	str1 := "42"
	str2 := " -042"
	str3 := "1337c0d3"
	str4 := "0-1"
	str5 := "words and 987"

	result1 := myAtoi(str1)
	result2 := myAtoi(str2)
	result3 := myAtoi(str3)
	result4 := myAtoi(str4)
	result5 := myAtoi(str5)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
	fmt.Println("Case 4:", result4)
	fmt.Println("Case 5:", result5)
}

func myAtoi(s string) int {
	result := 0
	index := 0
	multiplier := 1
	signal := 0
	digit := false

	runeMap := make(map[int]rune)

	for _, v := range s {

		if isDigit(v) {
			digit = true
			if index == 0 && v == 48 {
				continue
			} else {
				runeMap[index] = v
				index++
			}

		} else {
			if digit || signal != 0 {
				break
			}

			if !isWhitespace(v) {

				if isPositiveSign(v) {
					if signal != 0 {
						break
					}
					signal = 1

				} else if isNegativeSign(v) {
					if signal != 0 {
						break
					}
					signal = -1

				} else {
					break
				}
			}
		}
	}

	for i := index - 1; i >= 0; i-- {
		result = result + (int(runeMap[i]-48) * multiplier)
		multiplier *= 10

		if result > 2<<30-1 || multiplier/10 > 2<<30-1 {
			if signal == -1 {
				return -2 << 30
			} else {
				return 2<<30 - 1
			}
		}
	}

	if signal == -1 {
		result = result * -1
	}

	return result
}

func isDigit(d rune) bool {
	return d >= 48 && d <= 57
}

func isWhitespace(d rune) bool {
	return d == 32
}

func isNegativeSign(d rune) bool {
	return d == 45
}

func isPositiveSign(d rune) bool {
	return d == 43
}
