package main

import "fmt"

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func partition(s string) [][]string {
	var p, q int
	var i string
	p = 0
	q = 1
	for q - p <= len(s){
		i = s[p:q]
		if check(i){
			p++
			q++
		}else{

		}
		fmt.Println(i)
	}
	return [][]string{}
}

func check(i string) bool{
	for index := 0; index < len(i) / 2; index++{
		if i[index] != i[len(i) - 1 - index]{
			return false
		}
	}
	return true
}

func main() {
	s := "aab"
	partition(s)
}
