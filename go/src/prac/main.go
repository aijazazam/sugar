package main

import (
	"fmt"
	"time"
)

type ABC struct {
	c chan int
}

func main() {

	ch := make( chan int, 10 )
	st := make( chan ABC, 10 )

	abc := ABC{}
	abc.c = make( chan int, 10 )

	ch <- 100
	st <- abc
	abc.c <- 500

	go GoCallme( ch, st )

	time.Sleep( 3 * time.Second )
}

func GoCallme( ch chan int, st chan ABC ) {

	for {
		select {
			case c := <-ch: {
				fmt.Println("Recieved ", c)
			}
			case s := <-st: {
				fmt.Println(s)
				ch = s.c
			}
		}
	}
}

