package main

import (
	"fmt"
	"time"
)

type Producer struct {
	ch *chan int
}

type Consumer struct {
	ch *chan int
}

func (p *Producer) produce(n int, m int) {
	for t := 0; t < n; t++ {
		// p.ch <- rand.Int()
		*p.ch <- t + m*10
		fmt.Println("procuce :", t+m*10, "cur len:", len(*p.ch))
		time.Sleep(200 * time.Millisecond)
	}
}

func (c *Consumer) consume() {
	for {
		fmt.Println(<-*c.ch)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	products := make(chan int)
	// products := make(chan int, 10)
	for i := 0; i < 1; i++ {
		p := Producer{ch: &products}

		go p.produce(10, i)

	}
	for i := 0; i < 5; i++ {
		c := Consumer{ch: &products}
		go c.consume()
	}
	wait := make(chan struct{})
	<-wait
}
