package leetcode

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	target := 123
	result := ReverseInteger(target)
	expectation := 321

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = -123
	result = ReverseInteger(target)
	expectation = -321

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = 120
	result = ReverseInteger(target)
	expectation = 21

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = 1534236469
	result = ReverseInteger(target)
	expectation = 0

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}
}
