package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true

func isValid(s string) bool {
	if len(s) == 0{
		return true
	}else if len(s) % 2 != 0{
		return false
	}
	var backets []string
	var backetsMap map[string]string
	backetsMap = map[string]string{
		")":"(",
		"]":"[",
		"}":"{",
	}
	for _, char := range s{
		backets = append(backets, string(char))
		if len(backets) < 2{
			//
		}else {
			if backetsMap[string(char)] == backets[len(backets)-2]{
				backets = backets[0:len(backets)-2]
			}
		}
	}
	if len(backets) != 0{
		return false
	}
	//var backetMap map[string]int
	//backetMap = make(map[string]int)
	//for index := range s{
	//	backetMap[string(s[index])] = index
	//	backetMap[string(s[index]) + "count"]++
	//	b1, ok1 := backetMap["("]
	//	b2, ok2 := backetMap["["]
	//	b3, ok3 := backetMap["{"]
	//	if string(s[index]) == ")" && (!ok1 || (index-b1)%2==0){
	//		return false
	//	}else if string(s[index]) == "]" && (!ok2 || (index-b2)%2==0){
	//		return false
	//	}else if string(s[index]) == "}" && (!ok3 || (index-b3)%2==0){
	//		return false
	//	}
	//}
	//if backetMap["(count"] != backetMap[")count"] || backetMap["[count"] != backetMap["]count"] || backetMap["{count"] != backetMap["}count"]{
	//	return false
	//}
	return true
}

func main() {
	fmt.Println(isValid("}{"))
}