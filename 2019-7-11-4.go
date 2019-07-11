package main

import "fmt"

/*
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

示例 1:

输入: a = 1, b = 2
输出: 3
示例 2:

输入: a = -2, b = 3
输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func getSum(a int, b int) int {
	var sum, carry int
	sum = 0
	for b != 0{
		sum = a^b
		carry = (a&b) << 1
		a = sum
		b = carry
	}
	return sum
}

func main() {
	fmt.Println(getSum(-0, 3))
}
