package main

import "fmt"

func main(){
	var a,b int
	var  x string
	fmt.Print("press 1 for Addition\n")
	fmt.Print("press 2 for Subtraction\n")
	fmt.Print("press 3 for Multiplication\n")
	fmt.Print("press 4 for Division\n")
	fmt.Print("-------------------------------\n")
	fmt.Println("Enter your Choice")
	fmt.Scan(&x)
	fmt.Println("Enter two number")
	fmt.Scan(&a,&b)
	switch x {
	case "1": 
	c :=a+b
	fmt.Printf("Result is =%d",c)
	
	case "2":
	c :=a-b
	fmt.Printf("Rwsult is = %d",c)
	
	case "3":
	c :=a*b
	fmt.Printf("Result is =%d",c)
	
	case "4":
	c :=a/b
	fmt.Printf("Result is =%d",c)
	
	default:{

	
		fmt.Printf("wrong input")


	}


}
}