package main

import (
    "fmt"
)

func main(){
    numbers1 := []int {2,7,11,15}
    numbers2 := []int {3,2,4}
    numbers3 := []int {3,3,3,3,3,3}
    target1 := 9
    target2 := 6
    target3 := 6
    expected1 := "[0,1]"
    expected2 := "[1,2]"
    expected3 := "[0,1]"

    solution1 := twoSum(numbers1, target1)
    solution2 := twoSum(numbers2, target2)
    solution3 := twoSum(numbers3, target3)

    fmt.Println("Case 1")
    fmt.Println("Input")
    fmt.Print("Nums= ")
    fmt.Println(numbers1)
    fmt.Print("Target= ")
    fmt.Println(target1)
    fmt.Println("Ouput")
    fmt.Println(solution1)
    fmt.Println("Expected")
    fmt.Println(expected1)

    fmt.Println("Case 2")
    fmt.Println("Input")
    fmt.Print("Nums= ")
    fmt.Println(numbers2)
    fmt.Print("Target= ")
    fmt.Println(target2)
    fmt.Println("Ouput")
    fmt.Println(solution2)
    fmt.Println("Expected")
    fmt.Println(expected2)

    fmt.Println("Case 3")
    fmt.Println("Input")
    fmt.Print("Nums= ")
    fmt.Println(numbers3)
    fmt.Print("Target= ")
    fmt.Println(target3)
    fmt.Println("Ouput")
    fmt.Println(solution3)
    fmt.Println("Expected")
    fmt.Println(expected3)
}

func twoSum(nums []int, target int) []int {
    
    solution := []int {}

    for i := 0; i < len(nums)-1; i++ {
        for n := i+1; n < len(nums); n++ {
            if n > i && !contains(solution, i) && !contains(solution, n) && nums[i] + nums[n] == target  {
                solution = append(solution, i, n)
            }
        }
    }

    return solution
}

func contains(arr []int, t int) bool {
    for i := 0; i < len(arr); i++ {
        if arr[i] == t{
            return true
        }
    }
    return false
}