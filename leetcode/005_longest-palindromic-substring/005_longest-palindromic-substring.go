package main

import (
	"fmt"
)

/*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Constraints:
	1 <= s.length <= 1000
	s consist of only digits and English letters.
*/

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] {
				if len(result) < len(s[i:j+1]) && isPalindrome(s[i:j+1]) {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}

func main() {
	str1 := "babad"
	str2 := "cbbd"

	result1 := longestPalindrome(str1)
	result2 := longestPalindrome(str2)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}
