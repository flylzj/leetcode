package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	flag := 0
	last := 0
	median := []int{0, 0}
	
	for i < len(nums1) || j < len(nums2) {
		median[0] = last
		if i == len(nums1){
			last = nums2[j]
			if flag == (len(nums1) + len(nums2)) / 2{
				median[1] = nums2[j]
				break
			}
			j++
			flag++
			continue
		}
		if j == len(nums2){
			last = nums1[i]
			if flag == (len(nums1) + len(nums2)) / 2{
				median[1] = nums1[i]
				break
			}
			i++
			flag++
			continue
		}
		if nums1[i] < nums2[j]{
			last = nums1[i]
			if flag == (len(nums1) + len(nums2)) / 2{
				median[1] = nums1[i]
				break
			}
			i++
		}else{
			last = nums2[j]
			if flag == (len(nums1) + len(nums2)) / 2{
				median[1] = nums2[j]
				break
			}
			j++
		}
		flag++
	}
	if (len(nums1) + len(nums2)) % 2 != 0{
		return float64(median[1])
	}else{
		return float64(median[0] + median[1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3}))
}