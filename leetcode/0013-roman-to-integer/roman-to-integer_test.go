package leetcode

import "testing"

func TestRomanToInteger(t *testing.T)  {
	target := "III"
	result := BestFunc(target)
	expectation := 3
	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = "IV"
	result = BestFunc(target)
	expectation = 4
	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = "IX"
	result = RomanToInteger(target)
	expectation = 9
	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = "LVIII"
	result = RomanToInteger(target)
	expectation = 58
	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	target = "MCMXCIV"
	result = BestFunc(target)
	expectation = 1994
	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

}
