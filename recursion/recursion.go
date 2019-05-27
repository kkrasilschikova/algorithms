package recursion

import (
	"strings"
)

/*
ListSum calculates sum of a list of numbers
*/
func ListSum(list []int) int {
	sum := 0

	if len(list) == 0 {
		return sum
	} else if len(list) == 1 {
		return list[0]
	} else {
		return list[0] + ListSum(list[1:])
	}
}

/*
ReverseString returns a new string that is a reverse of an input string
*/
func ReverseString(str string) string {
	input := []rune(strings.ReplaceAll(str, " ", ""))

	if len(input) == 0 {
		return ""
	} else if len(input) == 1 {
		return str
	} else {
		return string(input[len(input)-1]) + ReverseString(string(input[0:len(input)-1]))
	}
}
