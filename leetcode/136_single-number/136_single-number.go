package main

import (
	"fmt"
)

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1

Constraints:
	1 <= nums.length <= 3 * 104
	-3 * 104 <= nums[i] <= 3 * 104
	Each element in the array appears twice except for one element which appears only once.
*/

func singleNumber(nums []int) int {
	intMap := make(map[int]int, 0)

	for _, n := range nums {
		intMap[n]++
	}

	for i, v := range intMap {
		if v == 1 {
			return i
		}
	}
	return -1
}

func main() {
	numbers1 := []int{2, 2, 1}
	numbers2 := []int{4, 1, 2, 1, 2}
	numbers3 := []int{1}

	result1 := singleNumber(numbers1)
	result2 := singleNumber(numbers2)
	result3 := singleNumber(numbers3)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
}
