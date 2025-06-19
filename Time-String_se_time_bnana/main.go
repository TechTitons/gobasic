package main

import (
	"fmt"
	"time"
)

func main() {
	stringTime := "2024-03-15"
	Value, _ := time.Parse("2006-01-02",stringTime)
	fmt.Println("String to Time",Value)

}