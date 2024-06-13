package main

import (
	"fmt"
	"strconv"
)

func main() {

	i1 := 121
	i2 := -121
	i3 := 10
	i4 := 12345654321

	pal1 := isPalindrome(i1)
	pal2 := isPalindrome(i2)
	pal3 := isPalindrome(i3)
	pal4 := isPalindrome(i4)

	fmt.Println(pal1)
	fmt.Println(pal2)
	fmt.Println(pal3)
	fmt.Println(pal4)
}

func isPalindrome(x int) bool {

	str := strconv.Itoa(x)
	test := []rune(str)

	numberSize := len(str)

	if x < 0 {
		return false
	}

	for i := 0; i < numberSize/2; i++ {
		if test[i] != test[numberSize-i-1] {
			return false
		}
	}

	return true
}
