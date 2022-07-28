package concurency

import (
	"fmt"
	"time"
)

type ball struct {
	hits int32
}

func StartPingPong() {
	table := make(chan *ball)

	go pingPong("ping", table)
	go pingPong("pong", table)
	//time.Sleep(1 * time.Millisecond)
	table <- new(ball)
	time.Sleep(1 * time.Second)
	<-table

}

func pingPong(name string, table chan *ball) {
	for {
		Ball := <-table
		Ball.hits++
		fmt.Println(name," ",  Ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- Ball
	}
}
