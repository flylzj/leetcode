package main

import "fmt"

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//
//
//示例 1：
//
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
//示例 2：
//
//输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
//示例 3：
//
//输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。
//
//
//
//
//进阶：
//
//你能用 O(1)（即，常量）内存解决此问题吗？

//Definition for singly-linked list.

type ListNode struct {
Val int
Next *ListNode
}


func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
	var p, q *ListNode
	p = head
	q = head
	p = p.Next
	q = q.Next
	if q == nil{
		return false
	}
	q = q.Next
	for p != nil && q != nil{
		if p == q{
			return true
		}
		p = p.Next
		q = q.Next
		if q == nil{
			return false
		}
		q = q.Next
	}
	return false
}

func main() {
	l1 := &ListNode{3, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{0, nil}
	l4 := &ListNode{-4, nil}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2

	l5 := &ListNode{1, nil}
	l6 := &ListNode{2, nil}

	l5.Next = l6
	l6.Next = l5


	fmt.Println(hasCycle(l1))
	fmt.Println(hasCycle(l5))
	fmt.Println(hasCycle(&ListNode{1, nil}))
}