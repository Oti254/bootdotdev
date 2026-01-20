package main

import (
	"sync"
	"time"
)

/**
* sync.RWMutex - improves performance in read intensive processes
* multiple goroutines can safely read from the map simultaneously
* as many R.Lock() calls can occur at the same time
* only one goroutine can hold a Lock(), during which all R.Lock() calls are blocked
* maps are safe for concurrent read access
* not concurrent read/write or write/write
* allows all readers to access the map
* a single writer locks out all the other readers and writers
**/
type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

