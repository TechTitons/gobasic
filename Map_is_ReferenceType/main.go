package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["zeeshan"] = 10
	myMap["sibu"] = 20

	//Map ko ek aur variable me assign karein
	anotherMap := myMap
	//anotherMap main change karein
	anotherMap["mobin"] = 100
	//Origional map bhe change hoga
	fmt.Println("Original myMap\n",myMap)
	fmt.Println("anotherMap",anotherMap)
}