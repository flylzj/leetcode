package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	i := 0
	sp := ""
	minLen := 1
	for len(s) - i >= minLen{
		j := len(s)
		for j - i >= minLen{
			if isPalindrom(s[i:j]){
				sp = s[i:j]
				minLen = len(sp)
				break
			}
			j--
		}
		i++
	}
	return sp
}

func isPalindrom(s string) bool{
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	s2 := string(r)
	return s2 == s
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}
