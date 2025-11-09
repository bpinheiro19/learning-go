package main

import (
	"fmt"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows: string convert(string s, int numRows);

Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"

Example 3:
Input: s = "A", numRows = 1
Output: "A"

Constraints:
	1 <= s.length <= 1000
	s consists of English letters (lower-case and upper-case), ',' and '.'.
	1 <= numRows <= 1000
*/

func convert(s string, numRows int) string {

	result := ""
	stringList := make([]string, numRows)

	if numRows == 1 {
		return s
	}

	for i := 0; i < len(s); i += (2*numRows - 2) {
		index := 0

		for j := i; j < i+numRows && j < len(s); j++ {
			index = j % (2*numRows - 2)
			stringList[index] += s[j : j+1]
		}

		for k := i + numRows; k < len(s) && k < i+(2*numRows-2) && index > 1; k++ {
			index--
			stringList[index] += s[k : k+1]
		}

	}

	for _, s := range stringList {
		result += s
	}

	return result
}

func main() {
	str1, num1 := "PAYPALISHIRING", 3
	str2, num2 := "PAYPALISHIRING", 4
	str3, num3 := "A", 1

	result1 := convert(str1, num1)
	result2 := convert(str2, num2)
	result3 := convert(str3, num3)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
}
