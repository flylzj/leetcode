package main

import "fmt"

//实现 strStr() 函数。
//
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1:
//
//输入: haystack = "hello", needle = "ll"
//输出: 2
//示例 2:
//
//输入: haystack = "aaaaa", needle = "bba"
//输出: -1

func strStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}else if len(haystack) < len(needle) {
		return -1
	}
	for i:=0;i<len(haystack)-len(needle)+1;i++{
		if needle == haystack[i:len(needle)+i]{
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("mississippi", "issipi"))
	fmt.Println(strStr("mississippi", "pi"))
}
