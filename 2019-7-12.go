package main

import (
	"fmt"
)

/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-and-say
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


//递归版本 时间100% 空间 20%
//func countAndSay(n int) string {
//	if n <= 1{
//		return "1"
//	}else{
//		res := countAndSay(n - 1)
//		return count(res)
//	}
//}


// 非递归版本
func countAndSay(n int) string {
	var res string
	var i int
	res = "1"
	for i = 0; i < n - 1; i++{
		res = count(res)
	}
	return res
}

func count(s string) string {
	var res string
	var count int
	count = 0
	res = ""
	for index, i := range s{
		count++
		if index == len(s) - 1 || s[index + 1] != s[index]{
			res += string(count + 48) + string(i)
			count = 0
		}
	}
	return res
}

func main() {
	fmt.Println(countAndSay(5))
}