package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	time.Sleep(2*time.Second)
	time2:=time.Now()
	fmt.Println("Time 1 before time2",time1.Before(time2))
	fmt.Println("time2 after time1",time2.After(time1))
	fmt.Println("time1 and time2 are equal",time1.Equal(time2))
}