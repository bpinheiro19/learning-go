package main

import (
	"fmt"
)

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

Constraints:

	1 <= haystack.length, needle.length <= 104
	haystack and needle consist of only lowercase English characters.
*/
func main() {
	str1, str2 := "sad", "sad"
	str3, str4 := "leetcode", "leeto"

	result1 := strStr(str1, str2)
	result2 := strStr(str3, str4)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}

func strStr(haystack string, needle string) int {
	result := -1

	hSize := len(haystack)
	nSize := len(needle)

	for i := 0; i <= hSize-nSize; i++ {
		if haystack[i:i+nSize] == needle {
			return i
		}
	}
	return result
}
