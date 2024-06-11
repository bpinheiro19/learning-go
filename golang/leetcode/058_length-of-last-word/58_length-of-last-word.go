package main

import (
    "fmt"
    "strings"
)

func main(){
    
    string1 := "Hello World"
    output1 := 5

    string2 := "   fly me   to   the moon  "
    output2 := 4

    string3 := "luffy is still joyboy"
    output3 := 6

    fmt.Println("Case 1")
    fmt.Println(testOutput(string1, output1))

    fmt.Println("Case 2")
    fmt.Println(testOutput(string2, output2))

    fmt.Println("Case 3")
    fmt.Println(testOutput(string3, output3))
}

func lengthOfLastWord(s string) int {
    str := strings.Fields(s)
    return len(str[len(str)-1]);
}

func testOutput(s string, output int) bool{
    return lengthOfLastWord(s) == output
}