package main

import (
	"fmt"
)

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解


func maxSubArray(nums []int) int {
	res := 0
	b := 0
	for i:=0;i<len(nums);i++{
		if nums[i] + b > nums[i]{
			b += nums[i]
		}else{
			b = nums[i]
		}
		if i == 0{
			res = b
		}
		if b > res{
			res = b
		}
	}
	return res
}

func main(){
	fmt.Println(maxSubArray([]int{0}))
}