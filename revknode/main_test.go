package main

import (
	"fmt"
	"testing"
)

func TestRevKNode(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	root := reverseKGroup(list, 2)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
