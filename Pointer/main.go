package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int 
	//es line pr yellow line isleyai aaraha hai keuki app is line ko aur short kr skte hai eg ptr:=&num

	ptr = &num
	fmt.Println("The value of num",num)
	fmt.Println("The address of variable",&num)
	fmt.Println("The value of num using pointer",*ptr)
}