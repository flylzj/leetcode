package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func groupAnagrams(strs []string) [][]string {
	var res = make([][]string, 0)
	var strMap = make(map[string][]string)
	for _, str := range strs{
		key := getKey(str)
		if _, ok := strMap[key]; !ok{
			strMap[key] = []string{str}
		}else{
			strMap[key] = append(strMap[key], str)
		}
	}
	for _, v := range strMap{
		res = append(res, v)
	}
	return res
}

func getKey(s string) string{
	res := strings.Split(s, "")
	sort.Strings(res)
	return strings.Join(res, "")
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}