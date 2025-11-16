package main

import (
	"fmt"
)

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Example 1:
Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:
Input: num1 = "123", num2 = "456"
Output: "56088"

Constraints:
	1 <= num1.length, num2.length <= 200
	num1 and num2 consist of digits only.
	Both num1 and num2 do not contain any leading zero, except the number 0 itself.
*/

func reverseString(s string) string {
	b := []rune(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1, num2 = reverseString(num1), reverseString(num2)

	result := make([]rune, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		n1 := rune(num1[i] - 48)

		for j := 0; j < len(num2); j++ {
			n2 := rune(num2[j] - 48)

			mul := n1 * n2
			result[i+j] += mul
		}
	}

	bytes := make([]rune, 0)
	var carry rune = 0
	for _, e := range result {
		if carry > 0 {
			e += carry
			carry = 0
		}

		for e > 9 {
			e -= 10
			carry++
		}
		bytes = append(bytes, e+48)
	}

	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] != 48 {
			bytes = bytes[:i+1]
			break
		}
	}

	return reverseString(string(bytes))
}

func main() {
	nums1, nums2 := "2", "3"
	nums3, nums4 := "123", "456"

	result1 := multiply(nums1, nums2)
	result2 := multiply(nums3, nums4)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}
