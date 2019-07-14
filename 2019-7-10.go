package main

import "fmt"

type ListNode4 struct {
	Val int
	Next *ListNode4
}

//反转链表

func reverseList(head *ListNode4) *ListNode4 {
	if head == nil{
		return head
	}
	var p, q *ListNode4
	p = head
	q = head.Next
	for q != nil{
		PrintListNode(head)
		p.Next = q.Next
		q.Next = head
		head = q
		q = p.Next
	}
	return head
}

func PrintListNode(head *ListNode4){
	for head != nil{
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Print("\n")
}

func main() {
	head := reverseList(&ListNode4{
		1,
		&ListNode4{
			2,
			&ListNode4{
				3,
				&ListNode4{
					4,
					&ListNode4{
						5,
						nil,
					},
				},
			},
		},
	})
	fmt.Println(head.Next.Next.Next.Next)
}
