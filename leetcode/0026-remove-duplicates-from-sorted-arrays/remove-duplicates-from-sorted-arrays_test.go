package leetcode

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}

	result := removeDuplicates(nums)

	expectation := 5

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}

	nums1 := []int{0,0,1,1,1,2,2,3,3,4}

	result1 := removeDuplicatesBest(nums1)
	expectation1 := 5
	if result1 != expectation1 {
		t.Fatalf("expect is %d but result is %d, but fail",expectation, result)
	}
}
