package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("started")
	time.Sleep(10 *time.Second)
	fmt.Println("1 second after")
}