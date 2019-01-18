package main

import (
	"fmt"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
	var strMap map[byte]int
	var start []byte
	for i:=len(s);i>0;i--{
		for j:=0;j<len(s)+1-i;j++{
			start = []byte{}
			for k:=0;k<i;k++{
				start = append(start, s[k+j])
			}
			strMap = make(map[byte]int)
			for _, b := range start{
				if _, ok := strMap[b]; !ok{
					strMap[b] = 0
				}else{
					break
				}
			}
			if len(strMap) == len(start){
				return len(start)
			}
		}
	}
	return len(start)
}

func lengthOfLongestSubstring2(s string) int{
	var index [256]int
	index = [256]int{}
	n := len(s)
	var i, ans int
	i = 0
	ans = 0
	for j:= 0;j<n;j++{
		//i = int(math.Max(float64(index[s[j]]), float64(i)))
		//ans = int(math.Max(float64(ans), float64(j - i + 1)))
		if index[s[j]] > i {
			i = index[s[j]]
		}
		if j - i + 1 > ans{
			ans = j - i + 1
		}
		index[s[j]] = j + 1
	}
	return ans

}

func main(){
	fmt.Println(lengthOfLongestSubstring2("hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
}
