package main

import "fmt"

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type ListNode struct {
	Val int
	Next *ListNode
}


func sortList(head *ListNode) *ListNode {
	var t,p, q *ListNode
	t = head
	var length int
	length = 0
	for t != nil{
		length++
		t = t.Next
	}
	if length <= 1{
		return head
	}
	p = head
	q = head.Next
	for i := 0; i < length - 1; i++{
		for j := 0; j < length - i - 1; j++{
			if p.Val > q.Val{
				t := p.Val
				p.Val = q.Val
				q.Val = t
			}
			p = p.Next
			q = q.Next
		}
		p = head
		q = head.Next
	}
	return head
}

func main() {
	res := sortList(&ListNode{
		4,
		&ListNode{
			2,
			&ListNode{
				1,
				&ListNode{
					3,
					&ListNode{
						0,
						nil,
					},
				},
			},
		},
	})
	for res != nil{
		fmt.Println(res)
		res = res.Next
	}
}
