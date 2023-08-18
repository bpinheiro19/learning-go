package main

import (
	"fmt"
)

func main() {

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {

		if i < len(s)-1 && checkSub(string(s[i:i+2])) {
			sum += convertNumber(string(s[i+1])) - convertNumber(string(s[i]))
			i++
		} else {
			sum += convertNumber(string(s[i]))
		}

	}
	return sum
}

func checkSub(s string) bool {
	return s == "IV" || s == "IX" || s == "XL" || s == "XC" || s == "CD" || s == "CM"
}

func convertNumber(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}
