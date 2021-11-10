package leetcode

import (
	"strings"
)

func RomanToInteger(s string) (result int){
	length := len(s)
	if length <= 0 || length > 15 {
		return 0
	}
	symbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	symbolsPlus := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	for key,amount := range symbolsPlus {
		count := strings.Count(s,key)
		if count > 0 {
			s = strings.Replace(s,key,"",count)
			result += count * amount
		}
	}

	for _,item := range strings.Split(s, "") {
		result += symbols[item]
	}

	return result
}

type band struct {
	i int    // integer
	n string // numeral
}

func BestFunc(s string) int {
	result := 0

	var bands = []band{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _,b := range bands {
		for strings.HasPrefix(s, b.n) {
			result += b.i
			s = s[len(b.n):]
		}
	}
	return result
}

