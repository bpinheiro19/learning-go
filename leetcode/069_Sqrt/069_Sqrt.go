package main

import (
	"fmt"
)

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Example 1:
Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.

Example 2:
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

Constraints:
	0 <= x <= 231 - 1
*/

func main() {
	num1 := 4
	num2 := 8

	result1 := mySqrt(num1)
	result2 := mySqrt(num2)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}

func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return x
}
