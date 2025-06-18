package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time", currentTime)

	formatted := time.Now().Format("2006-01-05,Monday,15:01:02")
	fmt.Println("Formatted Time", formatted)
}
