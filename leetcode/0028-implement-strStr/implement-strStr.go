package leetcode

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	ln := len(needle)

	for i := 0; i <= len(haystack)-ln; i++ {
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}
