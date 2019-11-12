package main

import (
	"fmt"
)

func myAtoi(str string) int {
	validS := make([]uint8, 0)
	var n int32
	firstNumber := 0
	neg := 2
	for i, s := range str{
		// 遇到的第一个不是数字也不是空格和负号
		if s < 48 || s > 57{
			if s != 32 && s != 45{
				return 0
			}
		}else if s >= 48 && s <= 57{
			firstNumber = i
			break
		}else{
			return 0
		}
	}
	fmt.Println(firstNumber)
	// 数字的前一个不是负号
	if firstNumber == 0{

	}else if str[firstNumber- 1] != 45{
		return 0
	}else {
		neg = 0
	}
	for i := firstNumber; i < len(str); i++ {
		if str[i] < 48 || str[i] > 57{
			break
		}
		validS = append(validS, str[i])
	}
	n = 0
	j := 1
	fmt.Println(validS)
	for i := len(validS) - 1; i >= 0; i--{
		n += int32((int(validS[i]) - 48) * j)
		j *= 10
	}
	n -= int32(2 - neg) * n
	return int(n)
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}