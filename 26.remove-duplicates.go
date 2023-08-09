package main

import (
    "fmt"
)

func main(){
    
    nums1 := []int {1,1,2}
    nums2 := []int {0,0,1,1,1,2,2,3,3,4}

	expectedNumbers1 := []int {1,2}
	expectedNumbers2 := []int {0,1,2,3,4}

	output1 := 2
	output2 := 5

    fmt.Print("Case 1: ")
    fmt.Println(testOutput(nums1, expectedNumbers1, output1))

    fmt.Print("Case 2: ")
    fmt.Println(testOutput(nums2, expectedNumbers2, output2))

}

func removeDuplicates(nums []int) int {
	unique := 1
  for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1]{
			nums[unique] = nums[i+1]
			unique++
		}
	}
	return unique
}

func testOutput(nums []int, expectedNums[]int, output int) bool{

	k := removeDuplicates(nums)

	for i := 0; i < len(expectedNums); i++ {	
		if nums[i] != expectedNums[i] {
			return false
		}
	}

    return k == output
}

