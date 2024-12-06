package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var singleList *ListNode = lists[0]
	for i := 1; i < len(lists); i++ {
		singleList = mergeTwoList(singleList, lists[i])
	}
	return singleList
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	var newListNode = &ListNode{}
	var current *ListNode = newListNode
	if list1.Val < list2.Val {
		current.Val = list1.Val
		list1 = list1.Next
	} else {
		current.Val = list2.Val
		list2 = list2.Next
	}
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
			current = current.Next
			continue
		}
		current.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return newListNode
}
