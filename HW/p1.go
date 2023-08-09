package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var recievedCount uint64
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 9; i++ {
			c <- i
			fmt.Printf("g1 sent <%v>\n", i)
		}
		close(c)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range c {
			atomic.AddUint64(&recievedCount, uint64(1))
			fmt.Printf("g2 recieved <%v>\n", n)
		}
	}()
	wg.Wait()
	fmt.Println("Recieved counts:", recievedCount)
}
