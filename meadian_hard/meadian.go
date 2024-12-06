package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, n int
	var nums = make([]int, len(nums1)+len(nums2))
	target := len(nums) / 2
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[int(target)-1]+nums2[int(target)]) / 2
		}
		return float64(nums2[int(target)])
	}
	if len(nums2) == 0 {
		if len(nums1)%2 == 0 {
			return float64(nums1[int(target)-1]+nums1[int(target)]) / 2
		}
		return float64(nums1[int(target)])
	}
	for i < len(nums1) && j < len(nums2) {
		if n == target+1 {
			if len(nums)%2 == 0 {
				return float64(nums[int(target)-1]+nums[int(target)]) / 2
			}
			return float64(nums[int(target)])
		}
		if nums1[i] == nums2[j] {
			nums[n] = nums1[i]
			n++
			i++
			nums[n] = nums2[j]
			n++
			j++
			continue
		}
		if nums1[i] > nums2[j] {
			// nums = append(nums, nums2[j])
			nums[n] = nums2[j]
			j++
			n++
			continue
		}
		// nums = append(nums, nums1[i])
		nums[n] = nums1[i]
		i++
		n++
	}

	if i != len(nums1) {
		for i < len(nums1) {
			if n == target+1 {
				if len(nums)%2 == 0 {
					return float64(nums[int(target)-1]+nums[int(target)]) / 2
				}
				return float64(nums[int(target)])
			}
			nums[n] = nums1[i]
			i++
			n++
		}
	}
	if j != len(nums2) {
		// nums = append(nums, nums2[j:]...)
		for j < len(nums2) {
			if n == target+1 {
				if len(nums)%2 == 0 {
					return float64(nums[int(target)-1]+nums[int(target)]) / 2
				}
				return float64(nums[int(target)])
			}
			nums[n] = nums2[j]
			j++
			n++
		}
	}
	if len(nums)%2 == 0 {
		mediansIdx := len(nums) / 2
		return float64(nums[mediansIdx-1]+nums[mediansIdx]) / 2
	}
	mediansIdx := (len(nums) + 1) / 2
	return float64(nums[mediansIdx-1]+nums[mediansIdx]) / 2
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	var slice = []int{}
// 	lenTotal := len(nums1) + len(nums2)
// 	odd := func() bool {
// 		if lenTotal%2 == 0 {
// 			return false
// 		}
// 		return true
// 	}()
// 	medIdx := lenTotal / 2

// }

func main() {
	nums1 := []int{1, 3, 4, 5}
	nums2 := []int{2, 6, 7, 8}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
