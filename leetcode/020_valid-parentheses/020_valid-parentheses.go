package main

import (
	"fmt"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([])"
Output: true

Example 5:
Input: s = "([)]"
Output: false

Constraints:
	1 <= s.length <= 104
	s consists of parentheses only '()[]{}'.
*/

// ()[]{} 40 41 91 93 123 125
func isValid(s string) bool {
	r := make([]rune, 0)
	for _, e := range s {

		if e == '(' || e == '[' || e == '{' {
			r = append(r, e)

		} else if e == ')' || e == '}' || e == ']' {
			if len(r) == 0 {
				return false
			}
			v := r[len(r)-1]
			r = r[:len(r)-1]

			if (e == 41 && e != v+1) || ((e == 93 || e == 125) && e != v+2) {
				return false
			}
		}
	}

	if len(r) > 0 {
		return false
	}

	return true
}

func main() {
	str1 := "([)]"
	str2 := "()[]{}"

	result1 := isValid(str1)
	result2 := isValid(str2)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}
