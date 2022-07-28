package concurency

import (
	"sync"
	"time"
)

// sync package says : "Package sync provides basic synchronization primitives such as mutual exclusion locks.
// Other than the Once and WaitGroup types, most are intended for use by low-level library routines.
// Higher-level synchronization is better done via channels and communication"
func DeadLock() {
	lock1 := &sync.Mutex{}
	lock2 := &sync.Mutex{}
	go func() {
		lock1.Lock()
		time.Sleep(1 * time.Millisecond)
		lock2.Lock()
		println("first Thread")
		lock2.Unlock()
		lock1.Unlock()
		
	}()

	go func() {
		lock2.Lock()
		time.Sleep(1 * time.Millisecond)
		lock1.Lock()
		println("second Thread")
		lock1.Unlock()
		lock2.Unlock()
	}()
}

// sleeps are set to ensure deadlock will happen