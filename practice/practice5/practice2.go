package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	c := make(chan int, 5)

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
	start := time.Now()
	go func(){
		defer wg.Done()
		time.Sleep(time.Second*2)
		for r := range c{
			fmt.Println("g2 recieved ", r)
		}
	}()
	wg.Wait()
	time_spent := time.Since(start)
	fmt.Println("Time spent for receiving ", time_spent - (time.Second*2))
}