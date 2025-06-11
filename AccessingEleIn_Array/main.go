package main

import "fmt"

func main() {
	var a [4]string = [4]string{"Apple", "Banana", "Orange"}
	fmt.Println(a[1])
	a[1]="mango"/*Modifying Element*/
	fmt.Println(a[1])
	fmt.Println(a)
	fmt.Println(len(a))
}