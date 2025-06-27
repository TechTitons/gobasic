package main

import "fmt"

func main() {
	fmt.Println("what is your name ")
	var name string
	fmt.Scan(&name)
	fmt.Println("hello,Mr", name)
}
