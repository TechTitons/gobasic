package main

import "fmt"

func main() {
	fmt.Println("the addition of two number")
	result :=add(2,4)
	fmt.Println("sum is",result)
	 
}

func add(a int, b int) int {
	return a+b
}
