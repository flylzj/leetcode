package main

import "fmt"

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3
//type ListNode struct {
//	Val int
//	Next *ListNode
//}


func deleteDuplicates(head *ListNode) *ListNode {
	var p *ListNode
	p = head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val{
			if p.Next.Next != nil{
				p.Next = p.Next.Next
			}else{
				p.Next = nil
			}
		}else{
			p = p.Next
		}
	}
	return head
}

func main() {
	l := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, nil}}})
	for l!=nil{
		fmt.Println(l.Val)
		l = l.Next
	}
}