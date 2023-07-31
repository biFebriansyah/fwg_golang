package main

import (
	"fmt"
	"strconv"
)

func WaitGP() {
	fmt.Println("program start")

	charChan := make(chan string)
	digitChan := make(chan string)

	wg.Add(2)

	go getChar("hello worlds", charChan)
	go getDigit([]int{1, 2, 3, 4, 5, 6}, digitChan)

	fmt.Println(<-charChan)
	fmt.Println(<-digitChan)

	wg.Wait()
	fmt.Println("\nProgram stop")
}

func getChar(st string, c chan string) {
	var result string

	for _, c := range st {
		result += string(c)
	}

	c <- result
	wg.Done()
}

func getDigit(dt []int, c chan string) {
	var result string

	for _, d := range dt {
		result += strconv.Itoa(d)
	}

	c <- result
	wg.Done()
}
