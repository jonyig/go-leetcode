package leetcode

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	target := 121
	result := isPalindrome(target)
	if !result {
		t.Fatalf("expect is true but result is %d, but fail", target)
	}

	targetFail := 10
	resultFail := PalindromeNumber(targetFail)
	if resultFail {
		t.Fatalf("expect is false but result is %d, but fail", targetFail)
	}

	targetNegative := -10
	resultNegative := PalindromeNumber(targetNegative)
	if resultNegative {
		t.Fatalf("expect is false but result is %d, but fail", targetNegative)
	}

	targetOverMax := 2147483648
	resultOverMax := PalindromeNumber(targetOverMax)
	if resultOverMax {
		t.Fatalf("expect is false but result is %d, but fail", targetOverMax)
	}
}
