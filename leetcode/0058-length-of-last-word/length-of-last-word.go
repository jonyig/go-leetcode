package leetcode

func LengthOfLastWord(s string) int {
	i := len(s)
	var result int
	for i >= 1{
		if s[i-1:i] != " " {
			result++
		}
		if s[i-1:i] == " " && result >0 {
			break
		}
		i--
	}
	return result
}
