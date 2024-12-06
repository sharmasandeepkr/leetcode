package main

import (
	"fmt"
	"testing"
)

func TestMeadian(t *testing.T) {
	nums1 := []int{2, 2, 4, 4}
	nums2 := []int{2, 2, 2, 4, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
