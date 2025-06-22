package main

import "fmt"

func increment(num int) {
	num = num + 1
	fmt.Println("the value of num ",num)
}
func main() {
	var num int=10
	fmt.Println("the original value of num ",num)
	increment(num)

}