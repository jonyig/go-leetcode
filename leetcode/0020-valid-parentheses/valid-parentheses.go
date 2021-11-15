package leetcode

import (
	"strings"
)

func ValidParentheses(s string) bool {
	isHasReplace := true
	parentheses := []string{
		"()",
		"[]",
		"{}",
	}

	for isHasReplace {
		isHasReplace = false
		for _,item := range parentheses {
			if strings.Contains(s,item)  {
				s = strings.ReplaceAll(s,item,"")
				isHasReplace = true
				break
			}

		}
	}

	if s != "" {
		return false
	}
	return true
}
// leetcode best answer
func IsValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []rune

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
