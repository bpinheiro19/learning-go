package main

import (
	"fmt"
)

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Constraints:
	1 <= nums.length <= 104
	-104 <= nums[i] <= 104
	nums contains distinct values sorted in ascending order.
	-104 <= target <= 104
*/

func main() {
	nums := []int{1, 3, 5, 6}

	result1 := searchInsert(nums, 5)
	result2 := searchInsert(nums, 2)
	result3 := searchInsert(nums, 7)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {

		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid

		} else if nums[mid] < target {
			low = mid + 1

			if low > len(nums)-1 || target < nums[low] {
				return mid + 1
			}

		} else if nums[mid] > target {
			high = mid - 1

			if high < 0 || target > nums[high] {
				return mid
			}
		}
	}
	return -1
}
