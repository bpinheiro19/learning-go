package main

import (
	"fmt"
)

func main() {
	array1 := []int{2, 4, 3}
	array2 := []int{5, 6, 4}

	l1 := newListNode(array1)
	l2 := newListNode(array2)

	output1 := mergeTwoLists(&l1, &l2)
	printListNode(output1)
	fmt.Println("")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	newList := &ListNode{0, nil}
	head := newList

	for list1 != nil {
		head.Next = &ListNode{list1.Val, nil}
		head = head.Next
		list1 = list1.Next
	}

	head2 := newList

	for list2 != nil {

		for head2.Next != nil && head2.Next.Val < list2.Val {
			head2 = head2.Next
		}

		next := head2.Next
		head2.Next = &ListNode{list2.Val, nil}
		head2 = head2.Next
		head2.Next = next
		list2 = list2.Next
	}

	return newList.Next
}

func newListNode(array []int) ListNode {
	head := &ListNode{array[len(array)-1], nil}

	for i := len(array) - 2; i >= 0; i-- {

		ele := array[i]

		new := &ListNode{ele, nil}

		second := head
		head = new
		head.Next = second

	}
	return *head
}

func printListNode(l *ListNode) {

	fmt.Print(l.Val)

	if l.Next != nil {
		fmt.Print(" -> ")
		printListNode(l.Next)
	}
}
