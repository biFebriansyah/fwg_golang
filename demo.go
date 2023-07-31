package main

import (
	"fmt"
	"net/http"
)

func Demo() {
	website := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"https://medium.com/",
		"https://golang.org/",
		"https://www.fazztrack.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, url := range website {
		wg.Add(1)
		go CheckWeb(url)
	}

	wg.Wait()

}

func CheckWeb(url string) {
	defer wg.Done()

	if res, err := http.Get(url); err != nil {
		fmt.Println(url, "is OKE")
	} else {
		fmt.Printf("[%d] %s is up \n", res.StatusCode, url)
	}
}
