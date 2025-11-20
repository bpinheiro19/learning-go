package main

import (
	"fmt"
)

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:
  The number of nodes in the list is in the range [0, 300].
  -100 <= Node.val <= 100
  The list is guaranteed to be sorted in ascending order.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		deleteDuplicates(head.Next)
	}
	return head
}

func main() {
	head1 := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	head2 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}

	result1 := deleteDuplicates(head1)
	result2 := deleteDuplicates(head2)

	fmt.Print("Case 1: ")
	printListNode(result1)
	fmt.Println()

	fmt.Print("Case 2: ")
	printListNode(result2)
}

func printListNode(l *ListNode) {

	fmt.Print(l.Val, " ")

	if l.Next != nil {
		fmt.Print("-> ")
		printListNode(l.Next)
	}
}
