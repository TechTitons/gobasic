package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var ptr *int
	for i := 0; i < 5; i++ {
		ptr = &a[i]
		fmt.Printf("Address of a %d\n",ptr)
		
		


	}
}