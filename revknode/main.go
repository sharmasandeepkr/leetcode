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

func reverseKGroup(head *ListNode, k int) *ListNode {
	newListNode := &ListNode{}
	current := newListNode

	var idx int
	var start = true
	for {
		bufferKeys := []int{}
		for idx < k {
			if head == nil {
				for j := 0; j < idx; j++ {
					current.Next = &ListNode{
						Val: bufferKeys[j],
					}
					current = current.Next
				}
				return newListNode
				// flush remaining
			}
			bufferKeys = append(bufferKeys, head.Val)
			idx++
			head = head.Next
		}
		idx = 0
		if start {
			start = false
			i := k - 1
			current.Val = bufferKeys[i]
			i--

			for i := k - 2; i >= 0; i-- {
				current.Next = &ListNode{
					Val: bufferKeys[i],
				}
				current = current.Next
			}
			continue
		}
		for i := k - 1; i >= 0; i-- {
			current.Next = &ListNode{
				Val: bufferKeys[i],
			}
			current = current.Next
		}

	}

}
