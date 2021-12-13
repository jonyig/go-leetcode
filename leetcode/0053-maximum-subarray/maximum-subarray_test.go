package leetcode

import "testing"

func TestName(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	expectation := 6

	result := maxSubArray(nums)

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}
}
