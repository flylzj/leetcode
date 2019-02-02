package main

import "fmt"

//给定一个整数 n，返回 n! 结果尾数中零的数量。
//
//示例 1:
//
//输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。
//示例 2:
//
//输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零.

//   1  2  3  4  5  6  7  8  9  10  11  12
//   0  1  2  2  3  4  4  6  7  7   8   9

func trailingZeroes(n int) int {
	var count int
	count = -1
	for i:=0; i<=n; i+=5{
		if i % 5 == 0{
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(trailingZeroes(10))
}