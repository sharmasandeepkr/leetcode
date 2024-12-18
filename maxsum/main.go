package main

// func add2Stack(buffer []*int, n int, currentSum int) (sum int) {
// 	sum = currentSum - (*buffer[0]) + n
// 	buffer = append(buffer, &n)
// 	buffer = append(buffer[0:0], buffer[1:]...)
// 	return
// }

func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]struct{})

	for _, val := range banned {
		bannedMap[val] = struct{}{}
	}
	currentNO := 1
	buffer := []int{}
	nthIdx := 0
	bufferSum := 0
	for {
		if nthIdx == n {
			break
		}
		_, exist := bannedMap[currentNO]
		if exist {
			currentNO++
			continue
		}
		buffer = append(buffer, currentNO)
		bufferSum += currentNO
		currentNO++
		nthIdx++
		if bufferSum == maxSum {
			return nthIdx
		}
	}
	if bufferSum < maxSum {
		return n
	}
	for i := 1; i < len(buffer); i++ {
		newBufferWidth := i
		newBufferSum := 0
		for sidx := 0; sidx < len(buffer)-i; sidx++ {
			for j:=0; j< newBufferWidth; j++{
				newBufferSum+= 
			}
		}
	}
	return 0
}
