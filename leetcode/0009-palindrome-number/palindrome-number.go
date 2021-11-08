package leetcode

import (
	"strconv"
)

func PalindromeNumber(target int) bool {

	if target > 2147483648-1 || -2147483648 > target {
		return false
	}
	//log.Print(int32(target))
	if target < 0 {
		return false
	}

	reverse, _ := strconv.Atoi(ReverseString(strconv.Itoa(target)))

	if reverse == target {
		return true
	}

	return false
}

func ReverseString(str string) string {
	var output string
	for _, char := range str {
		output = string(char) + output
	}
	return output
}

// leetcode clean code
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	copy := x
	reverse := 0

	for x > 0 {
		reverse *= 10
		reverse += x % 10
		x /= 10
	}
	return reverse == copy
}
