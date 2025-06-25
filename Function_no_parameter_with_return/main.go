package main

import (
	
	"fmt"
	


)
func main(){
	s:=add()
	fmt.Println("sum is ",s)

}

func add() int {
	var num1,num2 int
	fmt.Println("enter first number")
	fmt.Scan(&num1)
	fmt.Println("enter second number")
	fmt.Scan(&num2)
	sum :=num1+num2
	
	
	return sum
	
	
}
