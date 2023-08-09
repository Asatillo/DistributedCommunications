package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	ints := []int{1, 2, 3, -4, 5}

	intMap := make(map[string]int)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range ints {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			mu.Lock()
			intMap[strconv.Itoa(a)] = handleInt(a)
			mu.Unlock()
		}(n)
	}

	wg.Wait()
	fmt.Println(intMap)
}

func handleInt(n int) int { return (n + 2) * 3 }
