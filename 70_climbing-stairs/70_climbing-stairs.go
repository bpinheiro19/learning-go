package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(45))
}

func climbStairs(n int) int {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		sum := x + y
		x = y
		y = sum
	}
	return y
}
