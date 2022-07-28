package concurency


import (
	"fmt"
	"time"
)

type token struct {
	tokenChan chan int
}

type petriNet struct {
	token       token
	readers     int
	shared_data int
}

func NewPn(readers int) *petriNet {
	t := token{make(chan int, readers)}
	for i := 0; i < readers; i++ {
		t.tokenChan <- 1
	}
	return &petriNet{token: t, readers: readers}
}

func (p *petriNet) startCycle() {
	p.startWriting()
	p.startReading()
}

func (p *petriNet) startReading() {
	go func(p *petriNet) {
		for i := 0; i < p.readers*2; i++ {
			go func(p *petriNet, interval int) {
				for {
					time.Sleep(10 * time.Millisecond)
					p.read()
				}
			}(p, i)
		}
	}(p)
}

func (p *petriNet) startWriting() {
	go func(p *petriNet) {
		i := 0
		for {
			time.Sleep(1 * time.Millisecond)
			p.write(i)
			i++
		}
	}(p)
}

func (p *petriNet) read() int {
	<-p.token.tokenChan
	defer func() {
		fmt.Println("p.read(): ", p.shared_data, " ", time.Now(), " ", len(p.token.tokenChan))
		p.token.tokenChan <- 1
	}()

	return p.shared_data
}

func (p *petriNet) write(newVal int) {
	fmt.Println("p.write() started: ", p.shared_data, " ", time.Now(), " ", len(p.token.tokenChan))
	for i := 0; i < p.readers; i++ {
		<-p.token.tokenChan
	}
	defer func() {
		fmt.Println("p.write() ended: ", p.shared_data, " ", time.Now(), " ", len(p.token.tokenChan))
		for i := 0; i < p.readers; i++ {
			p.token.tokenChan <- 1
		}
	}()
	p.shared_data = newVal
}

// go.dev says : Do not communicate by sharing memory; instead, share memory by communicating.
func (p *petriNet) StartConcurrentRead(newVal int) {
	Pn := NewPn(2)
	Pn.startCycle()
}

