package main

import (
	"fmt"
	"sort"
)

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:
	nums1.length == m
	nums2.length == n
	0 <= m <= 1000
	0 <= n <= 1000
	1 <= m + n <= 2000
	-106 <= nums1[i], nums2[i] <= 106
*/

func main() {
	nums1, nums2 := []int{1, 3}, []int{2}
	nums3, nums4 := []int{1, 2}, []int{3, 4}

	result1 := findMedianSortedArrays(nums1, nums2)
	result2 := findMedianSortedArrays(nums3, nums4)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mergedNum := append(nums1, nums2...)

	sort.Slice(mergedNum, func(i, j int) bool {
		return mergedNum[i] < mergedNum[j]
	})

	result := 0.0
	size := len(mergedNum)

	if size%2 == 0 {
		result = (float64(mergedNum[size/2]) + float64(mergedNum[size/2-1])) / 2
	} else {
		result = float64(mergedNum[size/2])
	}

	return result
}
