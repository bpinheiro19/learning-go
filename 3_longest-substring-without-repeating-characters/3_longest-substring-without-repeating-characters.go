package main

import "fmt"

func main() {

	string1 := "dvdf" //"davdfcb"
	string2 := "bbbbb"
	string3 := "pwwkew"
	string4 := " "

	output1 := 3
	output2 := 1
	output3 := 3
	output4 := 1

	result1 := lengthOfLongestSubstring(string1)

	fmt.Println("Result", result1)
	fmt.Println("Output", output1)

	result2 := lengthOfLongestSubstring(string2)

	fmt.Println("Result", result2)
	fmt.Println("Output", output2)

	result3 := lengthOfLongestSubstring(string3)

	fmt.Println("Result", result3)
	fmt.Println("Output", output3)

	result4 := lengthOfLongestSubstring(string4)

	fmt.Println("Result", result4)
	fmt.Println("Output", output4)
}

func lengthOfLongestSubstring(s string) int {

	test := []rune(s)
	usedLetters := make(map[string]int)

	x, y := 0, 0
	maxSize := 0
	currSize := 0

	for x < len(s) && y < len(s) {

		if usedLetters[string(test[y])] == 0 {
			currSize += 1
			usedLetters[string(test[y])] = 1
			y++

		} else {
			if maxSize < currSize {
				maxSize = currSize
			}

			currSize = 0
			usedLetters = make(map[string]int)
			x++
			y = x
		}

	}

	if maxSize < currSize {
		maxSize = currSize
	}

	return maxSize
}
