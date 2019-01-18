package main

import "fmt"

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}else if l2 == nil{
		return l1
	}
	var res, p *ListNode
	res = &ListNode{}
	p = res
	for{

		if l1.Val > l2.Val{
			res.Val = l2.Val
			l2 = l2.Next
		}else{
			res.Val = l1.Val
			l1 = l1.Next
		}
		if l1 == nil{
			res.Next = l2
			break
		}else if l2 == nil{
			res.Next = l1
			break
		}
		res.Next = &ListNode{}
		res = res.Next
	}
	return p
	//var p, q *ListNode
	//p = l1
	//q = l1
	//if l1 == nil && l2 == nil{
	//	return nil
	//}else if l1 == nil && l2 != nil{
	//	l1 = l2
	//	p = l1
	//	q = l1
	//}else{
	//	for l1.Next != nil{
	//		l1 = l1.Next
	//	}
	//	l1.Next = l2
	//}
	//for{
	//	flag := 0
	//	for p.Next != nil{
	//		if p.Val > p.Next.Val{
	//			tmp := p.Val
	//			p.Val = p.Next.Val
	//			p.Next.Val = tmp
	//			flag += 1
	//		}
	//		p = p.Next
	//	}
	//	if flag == 0{
	//		break
	//	}
	//	p = q
	//}
	//return q
}

func mainList(nums []int)*ListNode{
	var l,p *ListNode
	l = &ListNode{}
	p = l
	for index, val := range nums{
		l.Val = val
		if index < len(nums)-1{
			l.Next = &ListNode{}
			l = l.Next
		}
	}
	return p
}

func main() {
	l1 := mainList([]int{1,2,3})
	l2 := mainList([]int{2,3,4})
	l3 := mergeTwoLists(l1, l2)
	for l3!=nil{
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}