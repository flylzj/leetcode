package main

import (
	"fmt"
	)

//给定一个正整数，返回它在 Excel 表中相对应的列名称。
//
//例如，
//
//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB
//...
//示例 1:
//
//输入: 1
//输出: "A"
//示例 2:
//
//输入: 28
//输出: "AB"
//示例 3:
//
//输入: 701
//输出: "ZY"


func convertToTitle(n int) string {
	var res string
	var rev string
	var mod int
	for n > 0{
		mod = n % 26
		n /= 26
		if mod == 0{
			mod = 26
			n -= 1
		}
		rev += string(byte(mod + 64))

	}
	for i:=len(rev)-1; i >=0; i--{
		res += string(byte(rev[i]))
	}
	return res
}

func main() {
	fmt.Println(convertToTitle(1335))
}

