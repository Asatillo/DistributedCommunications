package main

import (
	"fmt"
	"time"
)

func main(){
	c:=make(chan chan string)
	in:=make(chan string)

	go func(){
		insideC:=<-c
		fmt.Println("g1 receive the channel inside")
		insideC<-"Hello"
		fmt.Println("g1 sent Hello to the channel inside")
	}()

	select{
	case c<-in:
		fmt.Println("Main send the channel inside")
		resp:= <-in
		fmt.Println("Main received ", resp)
	case<- time.After(time.Second*3):
			fmt.Println("Timeout, Quit")
			break
	}
}