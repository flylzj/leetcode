package main

import "fmt"

//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。
//
//说明：一个单词是指由字母组成，但不包含任何空格的字符串。
//
//示例:
//
//输入: "Hello World"
//输出: 5

func lengthOfLastWord(s string) int {
	if len(s) == 0{
		return 0
	}
	res := 0
	start := false
	for i:=len(s)-1;i>=0;i--{
		if (65<=s[i] && s[i]<=90) || (97<=s[i] && s[i]<=122){
			res++
			start = true
		}else if start && s[i]==32{
			break
		}
	}
	return res
}

func main(){
	//fmt.Println(byte('a'), byte('A'), byte('z'), byte('Z'), byte(' '))
	fmt.Println(lengthOfLastWord("Hello World"))
}