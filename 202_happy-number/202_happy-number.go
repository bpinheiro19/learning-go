package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	num1 := 19
	num2 := 2

	result1 := isHappy(num1)
	result2 := isHappy(num2)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
}

func isHappy(n int) bool {

	m := make(map[int]int)

	for n != 1 {

		array := []rune(strconv.Itoa(n))
		n = 0

		for _, num := range array {
			i, err := strconv.Atoi(string(num))

			if err != nil {
				log.Fatal(err)
			}
			n += i * i
		}

		if isDuplicateNumber(m, n) {
			return false
		}
	}

	return true
}

// If duplicate then solution is not possible
func isDuplicateNumber(m map[int]int, n int) bool {
	fmt.Println(m)
	if m[n] == 1 {
		return true

	} else {
		m[n] = 1
	}
	return false
}
