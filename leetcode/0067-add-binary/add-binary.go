package leetcode

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a,b = b,a

	}

	for i := len(a) - len(b); i > 0; i-- {
		b = "0" + b
	}
	var result string
	var haveOne int
	for j := len(a); j > 0; j-- {
		sum := parseInt(a[j-1:j]) + parseInt(b[j-1:j]) + haveOne
		haveOne = 0
		if sum >= 2 {
			haveOne = 1
			sum = sum % 2
		}

		result = fmt.Sprintf("%d%s", sum, result)
	}

	if haveOne == 1 {
		result = fmt.Sprintf("%d%s", haveOne, result)
	}

	return result
}

func parseInt(target string) int {
	result, _ := strconv.Atoi(target)
	return result
}
