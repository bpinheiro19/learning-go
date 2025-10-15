package main

import (
	"fmt"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:
	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lowercase English letters if it is non-empty.
*/

func main() {
	list1 := []string{"flower", "flow", "flight"}

	list2 := []string{"dog", "racecar", "car"}

	result1 := longestCommonPrefix(list1)
	result2 := longestCommonPrefix(list2)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	for i := 1; i < len(strs); i++ {
		index := findLastestIndexInCommon(result, strs[i])

		if index == 0 {
			return ""
		} else {
			result = result[:index]
		}
	}
	return result
}

func findLastestIndexInCommon(str1 string, str2 string) int {

	size := len(str1)
	if size > len(str2) {
		size = len(str2)
	}

	result := 0
	for i := 0; i < size; i++ {

		if str1[i] != str2[i] {
			return result
		}
		result++
	}
	return result
}
