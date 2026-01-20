package main

import (
	"sync"
	"time"
)

/**
* mutex - 'mu'
* mutexes are used to lock access to data
* defer is used along mutexes to ensure we never forget to unlock the mutex
* maps aren't safe for concurrency if multiple goroutines are accessing same map
* the maps must be locked with a mutex
* mutex - mutual exclusion - excludes different threads from accessing the same data at the same time
* help avoid the concurrent read/write problem - one thread writing to a variable while another reads the same var
* panic!!! ensues
* only one thread can lock a mutex
* used when one wants to safely share memory resources
*
**/
type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	sc.slowIncrement(key)
	defer sc.mu.Unlock()
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

