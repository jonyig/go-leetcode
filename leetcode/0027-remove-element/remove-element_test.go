package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := removeElement(nums, val)
	expectation := 5

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}

	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	result1 := removeElement(nums1, val1)
	expectation1 := 2

	if result1 != expectation1 {
		t.Fatalf("expect is %d but result is %d, but fail", expectation1, result1)
	}

	nums = []int{1}
	val = 1
	result = removeElement(nums, val)
	expectation = 0

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}
}
