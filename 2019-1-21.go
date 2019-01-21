package main

import (
	"fmt"
	"math"
	"unicode"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false

func isPalindrome2(s string) bool {
	var preIndex, lastIndex int
	preIndex = 0
	lastIndex = len(s) - 1
	for lastIndex > 0 && preIndex <= len(s)/2{
		if !unicode.IsDigit(rune(s[preIndex])) && !unicode.IsLetter(rune(s[preIndex])){
			preIndex++
			continue
		}else if !unicode.IsDigit(rune(s[lastIndex])) && !unicode.IsLetter(rune(s[lastIndex])){
			lastIndex--
			continue
		}else if s[preIndex] == s[lastIndex] || math.Abs(float64(int(s[preIndex]) - int(s[lastIndex]))) == 32{
			preIndex++
			lastIndex--
		}else{
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome2("\"OP\""))
}