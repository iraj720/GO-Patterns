package main

import (
	concurency "test3/concurrency"
	"time"
)

type x struct{}

func (x *x) hello(i int) bool {
	return true
}

func main() {
	concurency.DeadLock()
	time.Sleep(5 * time.Second)
}
