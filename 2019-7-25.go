package main

import "fmt"

/*
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func topKFrequent(nums []int, k int) []int {
	var nMap map[int]int
	var res []int
	nMap = make(map[int]int)

	for _, n := range nums{
		if _, ok := nMap[n]; !ok{
			nMap[n] = 1
		}else{
			nMap[n]++
		}
	}
	for key, v := range nMap{
		if v >= k{
			res = append(res, key)
		}
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}
