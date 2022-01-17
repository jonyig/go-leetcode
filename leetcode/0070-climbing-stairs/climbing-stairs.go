package leetcode

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	result := 0
	prev := 1
	next := 2

	for i := 2; i < n; i++ {
		result = prev + next
		prev = next
		next = result
	}
	return result
}
