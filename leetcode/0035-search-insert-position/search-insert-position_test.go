package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {

	nums := []int{1,3,5,6}
	target := 5

	result := searchInsert(nums,target)
	expectation := 2

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}

	nums = []int{1,3,5,6}
	target = 2

	result = searchInsert(nums,target)
	expectation = 1

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}

	nums = []int{1,3,5,6}
	target = 7

	result = searchInsert(nums,target)
	expectation = 4

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}

	nums = []int{1,3,5,6}
	target = 0

	result = searchInsert(nums,target)
	expectation = 0

	if result != expectation {
		t.Fatalf("expect is %d but result is %d, but fail", expectation, result)
	}
}
