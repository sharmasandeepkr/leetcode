package main

import (
	"fmt"
	"testing"
)

func TestMaxSum(t *testing.T) {
	bannedBuffer := []int{1, 5, 6}
	fmt.Println(maxCount(bannedBuffer, 5, 6))
}
