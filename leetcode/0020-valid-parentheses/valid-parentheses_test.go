package leetcode

import "testing"

func TestValidParentheses(t *testing.T) {
	s := "()"
	expectation := true
	result := ValidParentheses(s)

	if !result {
		t.Fatalf("expect is %t but result is %t, but fail",expectation, result)
	}

	s = "()[]{}"
	expectation = true
	result = ValidParentheses(s)

	if !result {
		t.Fatalf("expect is %t but result is %t, but fail",expectation, result)
	}

	s = "(]"
	expectation = false
	result = ValidParentheses(s)

	if result {
		t.Fatalf("expect is %t but result is %t, but fail",expectation, result)
	}

	s = "{[]}"
	expectation = true
	result = ValidParentheses(s)

	if !result {
		t.Fatalf("expect is %t but result is %t, but fail",expectation, result)
	}

	s = "([])"
	expectation = true
	result = ValidParentheses(s)

	if !result {
		t.Fatalf("expect is %t but result is %t, but fail",expectation, result)
	}

	s = "(([]){})"
	expectation = true
	result = IsValid(s)

	if !result {
		t.Fatalf("expect is %t but result is %t, but fail",expectation, result)
	}
}
