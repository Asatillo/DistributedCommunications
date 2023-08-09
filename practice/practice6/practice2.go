package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	attach := make(chan request)
	reverse := make(chan request)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 500)
			s1 := strconv.Itoa(i)
			s2 := strconv.Itoa(i + 2)
			req := request{s1, s2, make(chan string)}
			attach <- req
			result := <-req.resp
			fmt.Println("(attach) g1 sent", req.first, req.second, "received", result)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 500)
			s1 := strconv.Itoa(i)
			s2 := strconv.Itoa(i + 2)
			req := request{s1, s2, make(chan string)}
			reverse <- req
			result := <-req.resp
			fmt.Println("(reverse) g2 sent", req.first, req.second, "received", result)
		}
	}()

L:
	for {
		select {
		case msg := <-attach:
			result := msg.first + msg.second
			msg.resp <- result
			fmt.Println("(attach) Main received", msg.first, msg.second, "sent result", result)
		case msg := <-reverse:
			result := msg.second + msg.first
			msg.resp <- result
			fmt.Println("(reverse) Main received", msg.first, msg.second, "sent result", result)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout, Quit")
			break L
		}
	}
}

type request struct {
	first  string
	second string
	resp   chan string
}
