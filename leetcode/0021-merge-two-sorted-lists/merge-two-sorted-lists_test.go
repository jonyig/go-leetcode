package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := &ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1,Next: &ListNode{Val: 3,Next: &ListNode{Val: 4}}}

	result := MergeTwoSortedLists(l1,l2)
	expectation := []int{1,1,2,3,4,4}

	var resultList []int
	for result != nil {
		resultList = append(resultList,result.Val)
		result = result.Next
	}

	if !reflect.DeepEqual(resultList, expectation) {
		t.Fatalf("expect is %#v but result is %#v, but fail", result, expectation)
	}

	l11 := &ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 4}}}
	l21 := &ListNode{Val: 1,Next: &ListNode{Val: 3,Next: &ListNode{Val: 4}}}
	resultUseIterative := UseIterativeMergeTwoSortedLists(l11,l21)
	expectationUseIterative := []int{1,1,2,3,4,4}

	var resultListUseIterative []int
	for resultUseIterative != nil {
		resultListUseIterative = append(resultListUseIterative,resultUseIterative.Val)
		resultUseIterative = resultUseIterative.Next
	}

	if !reflect.DeepEqual(resultListUseIterative, expectationUseIterative) {
		t.Fatalf("expect is %#v but result is %#v, but fail", result, expectation)
	}
}
