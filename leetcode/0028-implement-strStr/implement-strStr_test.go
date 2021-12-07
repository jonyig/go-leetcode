package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	haystack := "hello"
	needle := "ll"

	result := strStr(haystack,needle)
	expectation := 2

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}

	haystack = ""
	needle = ""

	result = strStr(haystack,needle)
	expectation = 0

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}

	haystack = "aaaaa"
	needle = "bba"

	result = strStr(haystack,needle)
	expectation = -1

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}
}
