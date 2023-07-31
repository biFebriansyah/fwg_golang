package main

import (
	"fmt"
	"sync"
)

var angka = 0
var mt = sync.RWMutex{}

func Mutexs() {

	for i := 0; i < 5; i++ {
		wg.Add(2)

		mt.RLock() // dont read
		go ShowAngka()

		mt.Lock() // dont write
		go AddAngka()
	}

	wg.Wait()
}

func ShowAngka() {
	fmt.Printf("angkanya %d \n", angka)
	mt.RUnlock()
	wg.Done()
}

func AddAngka() {
	defer func() {
		mt.Unlock()
		wg.Done()
	}()

	angka++
}
