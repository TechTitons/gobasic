package main
import "fmt"

func modifyMap(m map[string]int){
	m["zeeshan"]=100 //Original map me changes hoga
}
func main(){
	myMap:=make(map[string]int)
	myMap["zeeshan"]=10
	myMap["Sibu"]=20
	fmt.Println("Original",myMap)
	modifyMap(myMap)
	fmt.Println("After",myMap)
}