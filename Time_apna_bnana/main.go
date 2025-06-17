package main

import (
	"fmt"
	"time"
)

func main() {
	self_time_create := time.Date(2025,time.June,15,10,0,0,0,time.UTC)
	fmt.Println("Apna time bnaya maine",self_time_create)
}