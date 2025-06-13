package main

import ("fmt"
	   "strconv"
)
func main(){
	
	var decimal float64=12.34
	
	text :=strconv.FormatFloat(decimal, 'f',2,64)
	
	fmt.Println("Result",text)
}