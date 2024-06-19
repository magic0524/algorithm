package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenLink(head *ListNode) {
	printLink(head)
	odd := head
	even := head.Next
	evenHead := head.Next

	for odd != nil && even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = odd.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	printLink(head)
}

func printLink(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func reverLink(head *ListNode) *ListNode {
	printLink(head)
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		if tmp != nil {
			head = tmp
		}
	}
	printLink(pre)
	return head
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	slowCnt := 1
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		slowCnt++
	}

	if slowCnt%2 == 1 {
		slow = slow.Next
	}

	slow = reverLink(slow)
	printLink(head)

	for i := 0; i < slowCnt-1; i++ {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}

	return true
}
