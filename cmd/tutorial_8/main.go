package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	dbData  = []string{"id1", "id2", "id3", "id4", "id5"}
	wg      = sync.WaitGroup{}
	results = []string{}
	m       = sync.RWMutex{}
)

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total execution time %v \n", time.Since(t0))
	fmt.Printf("The results are %v", results)
}

func dbCall(i int) {
	// stimulate delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()

	wg.Done()
}

func save(result string) {
	m.Lock() // full lock
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() // only read lock
	fmt.Printf("The current results are %v \n", results)
	m.RUnlock()
}
