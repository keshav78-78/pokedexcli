package main

import (
	"fmt"
	"os"
)

func callbackExit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}
