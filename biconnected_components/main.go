package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type conResult struct {
	result [][]int
	locker *sync.RWMutex
}

type line struct {
	f int
	s int
}
type lineStorage struct {
	allLine map[line]struct{}
	locker  *sync.RWMutex
}

func criticalConnections(n int, connections [][]int) [][]int {
	st := time.Now()
	singleConnections := make(map[int]int)
	connMap := make(map[int][]int) // vertices with degree of connections
	for i := 0; i < len(connections); i++ {
		vertex1 := connections[i][0]
		vertex2 := connections[i][1]

		if v2, exist := connMap[vertex1]; !exist {
			connMap[vertex1] = []int{vertex2}
			singleConnections[vertex1] = vertex2
		} else {
			connMap[vertex1] = append(connMap[vertex1], vertex2)
			for j := 0; j < len(v2); j++ {
				if v2[j] == vertex2 {
					delete(singleConnections, vertex1)
				}
			}
		}
		if v1, exist := connMap[vertex2]; !exist {
			connMap[vertex2] = []int{vertex1}
			singleConnections[vertex2] = vertex1
		} else {
			connMap[vertex2] = append(connMap[vertex2], vertex1)
			for j := 0; j < len(v1); j++ {
				if v1[j] == vertex1 {
					delete(singleConnections, vertex2)
				}
			}
		}
	}

	cr := conResult{
		result: [][]int{},
		locker: &sync.RWMutex{},
	}
	for key, value := range singleConnections {
		// result = append(result, []int{key, value})

		tv, exist := singleConnections[value]
		if exist && tv == key {
			delete(singleConnections, value)
		}

	}
	fmt.Println("first break: ", time.Since(st))
	objchan := make(chan obj)
	wg := &sync.WaitGroup{}
	uline := lineStorage{
		locker:  &sync.RWMutex{},
		allLine: make(map[line]struct{}),
	}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go process(context.Background(), wg, connMap, &cr, objchan, &uline)
	}
	count := 0
	// dont use same connection again and again for  verfication
	for key, value := range connMap {
		for i := 0; i < len(value); i++ {
			objchan <- obj{
				key:   key,
				value: value[i],
			}
			count++
		}
	}
	fmt.Println("count: ", count)
	close(objchan)
	wg.Wait()
	fmt.Println("result: ", cr.result)
	fmt.Println("second break: ", time.Since(st))
	return cr.result
}

type obj struct {
	key   int
	value int
}

func process(ctx context.Context, wg *sync.WaitGroup, connMap map[int][]int, cr *conResult, objchan chan obj, uline *lineStorage) {
	defer wg.Done()
	for {
		var key int
		var value int
		var targetMap = make(map[int]bool)
		select {
		case <-ctx.Done():
			return
		case ob, ok := <-objchan:
			if !ok {
				return
			}
			key = ob.key
			value = ob.value
			targetMap[key] = false
			targetMap[value] = false
		}
		for k, v := range connMap {
			if k == key {
				continue
			}
			for i := 0; i < len(v); i++ {
				_, exist := targetMap[v[i]]
				if exist {
					targetMap[v[i]] = true
				}
			}
			if targetMap[key] && targetMap[value] {
				break
			}
			if !targetMap[key] || !targetMap[value] {
				targetMap[key] = false
				targetMap[value] = false
			}
		}

		if !targetMap[key] || !targetMap[value] {
			// optimised search
			uline.locker.Lock()
			_, exist := uline.allLine[line{
				f: key,
				s: value,
			}]
			if exist {
				uline.locker.Unlock()
				continue
			}
			_, exist = uline.allLine[line{
				f: value,
				s: key,
			}]
			if exist {
				uline.locker.Unlock()
				continue
			}
			uline.allLine[line{
				f: key,
				s: value,
			}] = struct{}{}
			cr.result = append(cr.result, []int{key, value})
			uline.locker.Unlock()

			//finally update result
			// cr.locker.Lock()
			// cr.locker.Unlock()
		}
	}
}
