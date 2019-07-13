package main

import (
	"fmt"
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
	// 在leetcode测试用例上运行结果是对的，不知道为什么运行就是错的
	//var words []int32
	//for _, i := range s{
	//	if (int(i) >= 48 && int(i) <= 57) || (int(i) >= 65 && int(i) <= 90) || (int(i) >= 97 && int(i) <= 122){
	//		words = append(words, i)
	//	}
	//}
	//for i := 0; i < len(words) / 2; i++{
	//	if words[i] != words[len(words) - 1 - i] &&
	//		words[i] != words[len(words) - 1 -i] - 32 &&
	//		words[i] - 32 != words[len(words) - 1 - i]{
	//		return false
	//	}
	//}
	//return true

	var pre, tail int
	pre = 0
	tail = len(s) - 1
	for pre <= tail && pre < len(s) && tail >= 0{
		if !unicode.IsDigit(rune(s[pre])) && !unicode.IsLetter(rune(s[pre])){
			pre++
			continue
		}

		if !unicode.IsDigit(rune(s[tail])) && !unicode.IsLetter(rune(s[tail])){
			tail--
			continue
		}
		//fmt.Println(string(s[pre]), string(s[tail]))

		if s[pre] == s[tail] ||
			(s[pre] == s[tail] - 32 && s[tail] >= 97)||
			(s[pre] - 32 == s[tail] && s[pre] >= 97){
			pre++
			tail--
		}else{
				return false
		}
	}
	return true

}

func main() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome2("1Q"))
	fmt.Println(isPalindrome2(""))
}