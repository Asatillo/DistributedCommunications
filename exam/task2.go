package main

import(
	"fmt"
	"sync/atomic"
	"sync"
)
/*
  The main function wait for g1 and g2 and print out ”Received count:” and receivedCount.
*/
func main(){
	var receivedCount uint64
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=0;i<9;i++{
			c <- i
			fmt.Println("g1 sent ", i)
		}
		close(c)
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		for r := range c{
			fmt.Println("g2 recieved ", r)
			atomic.AddUint64(&receivedCount, 1)
		}
	}()

	wg.Wait()

	fmt.Println("Received count: ", receivedCount)
}