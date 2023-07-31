package main

import "fmt"

func Chanels() {
	wg.Add(2)
	var ch = make(chan int)

	go rOne(ch)
	go rTwo(ch)

	// fmt.Println(<-ch)

	wg.Wait()
}

// routine rechiver
func rOne(ch chan int) {
	i := <-ch
	fmt.Println(i)

	wg.Done()
}

// routine sender
func rTwo(ch chan int) {
	defer func() {
		wg.Done()
	}()

	var tes int = 20

	ch <- tes
}
