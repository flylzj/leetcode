package main

import "fmt"

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

注意事项：您可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


// O(n)/O(n)
func firstUniqChar2(s string) int {
	var count map[int32]int
	count = make(map[int32]int)
	for _, i := range s{
		if _, ok := count[i]; !ok{
			count[i] = 0
		}
		count[i]++
	}
	for index, i := range s{
		if count[i] == 1{
			return index
		}
	}
	return -1
}

func firstUniqChar(s string) int {
	var nums [26]int32
	nums = [26]int32{0}
	for _, i := range s{
		nums[i - 'a']++
	}
	for index, i := range s{
		if nums[i - 'a'] == 1{
			return index
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
