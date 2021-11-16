package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//refs https://ithelp.ithome.com.tw/articles/10213265
//recursive solution
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if (l1 == nil) && (l2 == nil) {
		return nil

	} else if l1 == nil {
		return l2

	} else if l2 == nil {
		return l1
	}

	// General cases
	if l1.Val < l2.Val {
		l1.Next = MergeTwoSortedLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoSortedLists(l1, l2.Next)
		return l2
	}
}

//iterative solution
func UseIterativeMergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	var prev = l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	}

	if l2 != nil {
		prev.Next = l2
	}
	return l3.Next
}
