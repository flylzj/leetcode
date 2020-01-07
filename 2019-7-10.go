package main

import "fmt"

type ListNode6 struct {
	Val int
	Next *ListNode
}

//反转链表

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var p, q, tmp *ListNode
	p = head
	q = head.Next
	p.Next = nil
	for q.Next != nil{
		tmp = q.Next
		q.Next = p
		p = q
		q = tmp

	}
	q.Next = p
	head = q
	return head
}

func PrintListNode(head *ListNode){
	for head != nil{
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Print("\n")
}

func main() {
	head := reverseList(&ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	})
	PrintListNode(head)
}
