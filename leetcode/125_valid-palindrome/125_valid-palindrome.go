package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
	1 <= s.length <= 2 * 105
	s consists only of printable ASCII characters.
*/

func main() {
	str1 := "A man, a plan, a canal: Panama"
	str2 := "race a car"
	str3 := " "

	result1 := isPalindrome(str1)
	result2 := isPalindrome(str2)
	result3 := isPalindrome(str3)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
}

func isPalindrome(s string) bool {
	str := strings.ToLower(strings.Join(regexp.MustCompile(`[a-zA-Z0-9]*`).FindAllString(s, -1), ""))

	for low, high := 0, len(str)-1; low < high; low, high = low+1, high-1 {
		if str[low] != str[high] {
			return false
		}
	}
	return true
}
