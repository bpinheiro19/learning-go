package main

import (
	"fmt"
)

/*
Given two binary strings a and b, return their sum as a binary string.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

Constraints:
	1 <= a.length, b.length <= 104
	a and b consist only of '0' or '1' characters.
	Each string does not contain leading zeros except for the zero itself.
*/

func main() {
	a1 := "11"
	b1 := "1"

	a2 := "1010"
	b2 := "1011"

	result1 := addBinary(a1, b1)
	result2 := addBinary(a2, b2)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}

func addBinary(a string, b string) string {

	listA := []byte(a)
	listB := []byte(b)

	sizeA := len(listA) - 1
	sizeB := len(listB) - 1

	size := sizeA
	if size > sizeB {
		size = sizeB
	}

	leftover := make([]byte, 0)
	sizeDiff := 0

	if len(listA) > len(listB) {

		sizeDiff = sizeA - sizeB
		leftover = listA[:sizeDiff]
		listA = listA[sizeDiff:]

	} else if len(listA) < len(listB) {

		sizeDiff = sizeB - sizeA
		leftover = listB[:sizeDiff]
		listB = listB[sizeDiff:]
	}

	result := make([]byte, 0)

	carry := false

	for i := size; i >= 0; i-- {
		result, carry = sumBinary(result, listA[i], listB[i], carry)
	}

	if len(leftover) > 0 {
		for i := len(leftover) - 1; i >= 0; i-- {
			result, carry = sumBinary(result, leftover[i], 48, carry)
		}
	}

	if carry {
		result = append(result, 49)
	}

	str := ""

	for i := len(result) - 1; i >= 0; i-- {
		str += string(result[i])
	}

	return str

}

func sumBinary(result []byte, a byte, b byte, carry bool) ([]byte, bool) {

	if a == 49 {
		if a == b {
			// 1 + 1 = 0 (carry 1)
			if carry {
				result = append(result, 49)
			} else {
				result = append(result, 48)
				carry = true
			}

		} else {
			// 1 + 0 = 1
			if carry {
				result = append(result, 48)
			} else {
				result = append(result, 49)
			}
		}

	} else if a == 48 {
		if a == b {
			// 0 + 0 = 0
			if carry {
				result = append(result, 49)
				carry = false
			} else {
				result = append(result, 48)
			}

		} else {
			// 1 + 0 = 1
			if carry {
				result = append(result, 48)
			} else {
				result = append(result, 49)
			}

		}
	}
	return result, carry
}
