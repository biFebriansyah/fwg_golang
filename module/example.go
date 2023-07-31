package module

import (
	"fmt"
	"os"
)

// exported function
func Exampple() {
	defer fmt.Println("Helo From example1")
	fmt.Println("Helo From example2")
	defer fmt.Println("defer 1")
	os.Exit(25)
	fmt.Println("Helo From example3")
	fmt.Println("Helo From example4")
	defer fmt.Println("defer 3")
}

// unexported function
func testing() {
	fmt.Println("testing")
}
