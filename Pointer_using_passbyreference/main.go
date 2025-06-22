package main

import "fmt"

func increment(ptr *int) {
	*ptr+=1
	fmt.Println("the value of num ",*ptr)
}
func main() {
	var num int=10
	fmt.Println("the original value of num ",num)
	increment(&num)

}