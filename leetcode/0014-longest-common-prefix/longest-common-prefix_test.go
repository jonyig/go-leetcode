package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T)  {

	strList := []string{"flower","flow","flight"}
	expectation := "fl"
	result := LongestCommonPrefix(strList)

	if result != expectation {
		t.Fatalf("expect is %s but result is %s, but fail",expectation, result)
	}

	strList = []string{"dog","racecar","car"}
	expectation = ""
	result = LongestCommonPrefix(strList)

	if result != expectation {
		t.Fatalf("expect is %s but result is %s, but fail",expectation, result)
	}
}
