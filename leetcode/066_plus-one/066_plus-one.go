package main

import (
	"fmt"
)

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.

Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].

Constraints:
	1 <= digits.length <= 100
	0 <= digits[i] <= 9
	digits does not contain any leading 0's.
*/

func main() {
	digits1 := []int{1, 2, 3}
	digits2 := []int{4, 3, 2, 1}
	digits3 := []int{9}

	result1 := plusOne(digits1)
	result2 := plusOne(digits2)
	result3 := plusOne(digits3)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
}

func plusOne(digits []int) []int {
	index := len(digits) - 1

	if digits[index] == 9 {
		digits[index] = 0

		appendIndex := index - 1
		for i := index - 1; i >= 0; i-- {
			if digits[i] != 9 {
				break
			}
			appendIndex--
			digits[i] = 0
		}

		if appendIndex == -1 {
			digits = append([]int{1}, digits...)
		} else {
			digits[appendIndex] = digits[appendIndex] + 1
		}

	} else {
		digits[index] = digits[index] + 1
	}

	return digits
}
