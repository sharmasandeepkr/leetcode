package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

// [[0,1],[1,2],[2,0],[1,3],[3,4],[4,5],[5,3]]
func TestBiconn(t *testing.T) {
	// n := 6
	// buf := [][]int{}
	// buf = append(buf, []int{0, 1})
	// buf = append(buf, []int{1, 2})
	// buf = append(buf, []int{2, 0})
	// buf = append(buf, []int{1, 3})
	// buf = append(buf, []int{3, 4})
	// buf = append(buf, []int{4, 5})
	// buf = append(buf, []int{5, 3})
	// fmt.Println(criticalConnections(n, buf))

	n := 10000
	buf := [][]int{}
	file, err := os.Open("./data3")
	if err != nil {
		t.Fail()
	}
	defer file.Close()
	dataBytes, err := io.ReadAll(file)
	if err != nil {
		t.Fail()
	}

	for i := 0; i < len(dataBytes); i++ {
		if dataBytes[i] == '[' {
			i++
		}
		var (
			fs string
			ss string
		)
		for dataBytes[i] != ',' {
			fs += string(dataBytes[i])
			i++
		}
		if dataBytes[i] == ',' {
			i++
		}
		f, err := strconv.Atoi(fs)
		if err != nil {
			t.Fail()
		}
		for dataBytes[i] != ']' {
			ss += string(dataBytes[i])
			i++
		}
		if dataBytes[i] == ']' {
			i++
		}
		s, err := strconv.Atoi(ss)
		if err != nil {
			t.Fail()
		}
		buf = append(buf, []int{f, s})
	}
	fmt.Println(criticalConnections(n, buf))
}
