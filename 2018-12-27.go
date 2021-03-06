package main

import "fmt"

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807


 //type ListNode struct {
 //    Val int
 //    Next *ListNode
 //}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3, l4 *ListNode
    var flag int
    l3 = &ListNode{}
    flag = 0
    sum := 0
    for l1 != nil || l2 !=nil{
        if l1 == nil{
            l1 = &ListNode{}
        }
        if l2 == nil{
            l2 = &ListNode{}
        }
        if flag == 0{
            l4 = l3
        }
        flag ++
        sum = l1.Val + l2.Val + l3.Val
        l3.Val = 0
        l3.Val += sum % 10
        if sum >= 10{
            l3.Next = &ListNode{}
            l3.Next.Val = 1   // sum / 10
        }else if l1.Next == nil && l2.Next == nil && sum < 10{
            //
        }else {
            l3.Next = &ListNode{}
        }
        l3 = l3.Next
        l1 = l1.Next
        l2 = l2.Next
    }
    return l4
}


func main(){
	//l1 := &ListNode{Val:2, Next:&ListNode{Val:4, Next:&ListNode{Val:3,Next:nil}}}
	//l2 := &ListNode{Val:5, Next:&ListNode{Val:6, Next:&ListNode{Val:4,Next:nil}}}
	//l1 := &ListNode{Val:5, Next:nil}
	//l2 := &ListNode{Val:5, Next:nil}
    l1 := &ListNode{Val:9, Next:&ListNode{Val:9}}
    l2 := &ListNode{Val:1, Next:nil}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil{
	   fmt.Print(l3.Val)
		l3 = l3.Next
	}
}