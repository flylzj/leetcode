package main

import "fmt"

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//l := 0
	//p := head
	//for p != nil {
	//	l++
	//	p = p.Next
	//}
	//if l == n {
	//	head = head.Next
	//	return head
	//}
	//p = head
	//for l - n > 1 {
	//	p = p.Next
	//	l--
	//}
	//p.Next = p.Next.Next
	//return head
	// 1->2->3->4->5
	// i     i
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	p, q := dummy, dummy
	for i := 0; i <= n; i++ {
		q = q.Next
	}
	for q != nil{
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}

func printLinkList(head *ListNode){
	for head != nil{
		fmt.Print(head.Val, "->")
		head = head.Next
	}

}

func main() {
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	//head = ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	n := 2
	printLinkList(removeNthFromEnd(&head, n))
}
