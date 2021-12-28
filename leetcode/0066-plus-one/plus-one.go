package leetcode

func plusOne(digits []int) []int {
	lastPosition := len(digits) - 1
	var isNine bool
	for digits[lastPosition] == 9 {
		digits[lastPosition] = 0
		if lastPosition > 0 {
			lastPosition--
		}
		isNine = true
	}

	if isNine {
		if digits[0] == 0 {
			digits = append([]int{1}, digits...)
		}else {
			digits[0]++
		}
		return digits
	}
	digits[lastPosition]++
	return digits
}
