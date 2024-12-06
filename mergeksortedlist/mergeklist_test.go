package main

import (
	"fmt"
	"testing"
)

func TestMergeKList(t *testing.T) {
	lists := []*ListNode{}
	lists = append(lists, &ListNode{Val: 1, Next: &ListNode{Val: 15}})
	lists = append(lists, &ListNode{Val: 3})
	lists = append(lists, &ListNode{Val: 4})
	lists = append(lists, &ListNode{Val: 2})
	lists = append(lists, &ListNode{Val: 6})
	lists = append(lists, &ListNode{Val: 5})
	lists = append(lists, &ListNode{Val: 7})
	singleList := mergeKLists(lists)
	for singleList != nil {
		fmt.Println(singleList.Val)
		singleList = singleList.Next
	}
}
