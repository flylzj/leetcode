package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。

func longestCommonPrefix(strs []string) string {
	length := 0
	shortest := ""
	for index, str := range strs{
		if len(str) < length || index == 0{
			length = len(str)
			shortest = str
		}
	}
	var i int
	for i=0;i<length;i++{
		pre := ""
		for index, str := range strs{
			if index == 0{
				pre = string(str[i])
			}else if string(str[i]) != pre{
				return string(shortest[:i])
			}
		}
	}
	return string(shortest[:i])
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}