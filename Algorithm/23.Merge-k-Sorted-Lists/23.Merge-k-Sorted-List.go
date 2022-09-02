package mergeksortedlists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	for cur, cur1, cur2 := dummyHead, head1, head2; cur1 != nil || cur2 != nil; {
		if cur1 == nil {
			cur.Next = cur2
			break
		}
		if cur2 == nil {
			cur.Next = cur1
			break
		}

		if cur1.Val <= cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	if left > right {
		return nil
	}

	mid := left + (right-left)>>1
	return mergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	return merge(lists, 0, len(lists)-1)
}
