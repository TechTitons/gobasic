package main

import ("fmt"
	   "strconv"
)
func main(){
	var num1 int=123
	var text string="hello"
	result :=strconv.Itoa(num1)+text
	fmt.Println("Result",result)
}
//yaha pr output string me nhe aaiga keuki hm sirf operation perform kr rhe hai 