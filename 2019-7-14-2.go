package main

import "fmt"

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type ListNode5 struct {
	Val int
	Next *ListNode
}

func isPalindrome3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var slow, fast, pre, prepre, p1, p2 *ListNode
	slow = head
	fast = head.Next
	pre = nil
	prepre = nil
	for fast != nil && fast.Next != nil {
		//反转前半段链表
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
		//先移动指针再来反转
		pre.Next = prepre
		prepre = pre
	}
	p2 = slow.Next
	slow.Next = pre
	if fast == nil{
		p1 = slow.Next
	}else{
		p1 = slow
	}
	for p1 != nil {
		if p1.Val != p2.Val{
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func printListNode(head *ListNode){
	for head != nil{
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Print("\n")
}


func main() {
	var L1 ListNode
	L1 = ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				2,
				&ListNode{
					1,
					nil,
				},
			},
		},
	}
	fmt.Println(isPalindrome3(&L1))
}
