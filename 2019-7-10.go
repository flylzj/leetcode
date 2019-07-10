package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//反转链表

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	var p, q *ListNode
	p = head
	q = head.Next
	for q != nil{
		p.Next = q.Next
		q.Next = head
		head = q
		q = p.Next
	}
	return head
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
	fmt.Println(head.Next.Next.Next.Next)
}
