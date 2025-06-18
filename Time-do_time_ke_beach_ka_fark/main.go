package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	time.Sleep(2*time.Second)
	fark:=time.Since(currentTime)
	fmt.Println("fark kya hai ",fark)
}