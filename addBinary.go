package main

import (
	"fmt"
)

//给定两个二进制字符串，返回他们的和（用二进制表示）。
//
//输入为非空字符串且只包含数字 1 和 0。
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"

func addBinary(a string, b string) string {
	aIndex := len(a) - 1
	bIndex := len(b) - 1
	var res string
	var carry uint8
	carry = 0
	for aIndex >= 0 || bIndex >= 0{
		if aIndex < 0 && bIndex >= 0{
			if carry + b[bIndex] == 50{
				res += string(byte(48))
				carry = 1
			}else{
				res += string(byte(carry + b[bIndex]))
				carry = 0
			}
			bIndex--
			continue
		}else if bIndex < 0 && aIndex >= 0{
			if carry + a[aIndex] == 50{
				res += string(byte(48))
				carry = 1
			}else{
				res += string(byte(carry + a[aIndex]))
				carry = 0
			}
			aIndex--
			continue
		}
		if byte(a[aIndex] + b[bIndex] - 48 + carry) > 50{
			res += string(byte(48 + carry))
			carry = 1
		}else if  byte(a[aIndex] + b[bIndex] - 48 + carry) == 50 {
			res += string(byte(48))
			carry = 1
		}else{
			res += string(byte(a[aIndex] + b[bIndex] - 48 + carry))
			carry = 0
		}
		aIndex--
		bIndex--
	}
	if carry == 1{
		res += "1"
	}
	s := ""
	for i:=len(res)-1;i>=0;i--{
		s+= string(res[i])
	}
	return s
}

func main() {
	fmt.Println(byte('1'), byte('0'))
	fmt.Println(addBinary("110010", "10111"))
	//fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}