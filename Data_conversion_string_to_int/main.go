package main

import ("fmt"
	   "strconv"
)
func main(){
	
	var text string="123"
	
	result,_:=strconv.Atoi(text)
	
	fmt.Println("Result",result)
}
//is prohram me agar aap 123 daite hai to vo int me kr daiga mgar aap "abc" daite hai to result 0 aaiga