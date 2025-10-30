package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21

Constraints:
	-2**31 <= x <= 2**31 - 1
*/

func main() {
	nums1 := 123
	nums2 := -123
	nums3 := 120
	nums4 := 1534236469

	result1 := reverse(nums1)
	result2 := reverse(nums2)
	result3 := reverse(nums3)
	result4 := reverse(nums4)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
	fmt.Println("Case 4:", result4)
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	newStr := make([]byte, 0)

	startIndex := 0

	if str[0] == 45 {
		newStr = append(newStr, str[0])
		startIndex = 1
	}

	for i := len(str) - 1; i >= startIndex; i-- {
		newStr = append(newStr, str[i])
	}

	newStr = removeLeftZeros(newStr)

	result, err := strconv.Atoi(string(newStr))
	if err != nil {
		log.Fatal(err)
	}

	if float64(result) > math.Pow(2, 31)-1 || float64(result) < math.Pow(-2, 31) {
		result = 0
	}

	return result
}

func removeLeftZeros(bytes []byte) []byte {
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != 48 {
			bytes = append(bytes[i:])
			return bytes
		}
	}
	return bytes
}
