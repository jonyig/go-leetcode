package leetcode

func ReverseInteger(x int) int {
	var isNegative bool

	if x == 0 || isOut(x) {
		return 0
	}

	if x < 0 {
		x = -1 * x
		isNegative = true
	}
	reverse := 0

	for x > 0 {
		reverse *= 10
		reverse += x % 10
		x /= 10
	}

	if isNegative {
		reverse = reverse * -1
	}

	if isOut(reverse) {
		return 0
	}
	return reverse
}

func isOut(x int) bool {
	 if x > 2147483648-1 || -2147483648 > x {
	 	return true
	 }
	 return false
}
