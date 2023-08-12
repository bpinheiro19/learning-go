package main

import (
	"fmt"
)

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

func main() {
	array1 := []int{2, 4, 3}
	array2 := []int{5, 6, 4}

	l1 := newListNode(array1)
	l2 := newListNode(array2)

	array3 := []int{9, 9, 9, 9, 9, 9, 9}
	array4 := []int{9, 9, 9, 9}

	l3 := newListNode(array3)
	l4 := newListNode(array4)

	output1 := addTwoNumbers(&l1, &l2)
	printListNode(output1)
	fmt.Println("")

	output2 := addTwoNumbers(&l3, &l4)
	printListNode(output2)
	fmt.Println("")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	plusone := false

	listNode := &ListNode{0, nil}
	head := listNode

	for l1 != nil || l2 != nil || plusone {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if plusone {
			sum += 1
			plusone = false
		}

		if sum > 9 {
			sum -= 10
			plusone = true
		}

		head.Next = &ListNode{sum, nil}
		head = head.Next
	}

	return listNode.Next
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
