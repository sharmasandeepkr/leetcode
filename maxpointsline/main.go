package main

import (
	"math"
)

type generalLineFactor struct {
	slope    float64
	constant float64
	infx     int
}

func findGeneralLineFactor(yb, yi, xb, xi int) generalLineFactor {
	var m float64
	var linfx int = 0
	var c float64
	if xb-xi == 0 {
		m = math.Inf(1)
		linfx = xi
		c = math.Inf(-1)
	} else {
		m = float64(yb-yi) / float64(xb-xi)
		c = float64(yb) - (float64(xb) * m)
	}
	// if m == math.Inf(1) {
	// 	fmt.Println(linfx, m, c, yb, yi, xb, xi)
	// }
	return generalLineFactor{
		slope:    m,
		constant: c,
		infx:     linfx,
	}
}

type point struct {
	x int
	y int
}

func maxPoints(points [][]int) int {
	slopeMaps := make(map[generalLineFactor]map[point]point)
	maxSlopeCount := 0
	if len(points) <= 2 {
		return len(points)
	}
	for b := 0; b < len(points); b++ {
		yb := points[b][1]
		xb := points[b][0]
		for i := b; i < len(points); i++ {
			if yb == points[i][1] && xb == points[i][0] {
				continue
			}
			glv := findGeneralLineFactor(yb, points[i][1], xb, points[i][0])
			miniMap, exist := slopeMaps[glv]
			if !exist {
				minimap := make(map[point]point)
				minimap[point{x: points[i][0], y: points[i][1]}] = point{x: xb, y: yb}
				slopeMaps[glv] = minimap
				if maxSlopeCount == 0 {
					maxSlopeCount = 1
				}
				continue
			}

			_, pointExist := miniMap[point{x: points[i][0], y: points[i][1]}]
			if pointExist {
				continue
			}

			miniMap[point{x: points[i][0], y: points[i][1]}] = point{x: xb, y: yb}
			slopeMaps[glv] = miniMap
			if len(miniMap) > maxSlopeCount {
				maxSlopeCount = len(miniMap)
			}
		}
	}
	// for key, val := range slopeMaps {
	// 	if key.slope == math.Inf(1) {
	// 		fmt.Println(key, val, len(val))
	// 	}
	// 	if len(val) == 6 {
	// 		fmt.Println("final ", key, val)
	// 	}
	// }
	return maxSlopeCount + 1
}
