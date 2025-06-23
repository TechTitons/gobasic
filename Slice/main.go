package main

import "fmt"

func main() {
	array := [7]int{10, 20, 30, 40, 50, 60, 70}//ek array bnayai
	fmt.Println("Original Array",array)

	//ab isse array se ek slice bnate hai
   slice :=array[1:5]//20 30,40,50 starting lega mgar end ko nhe laiga
   fmt.Println("Initial Slice",slice)

   //an hm slice me element add krege
   slice =append(slice, 80,90)
   fmt.Println("after Adding",slice)

   //Element update krna
   slice[1]=999 //index 1 (30) ko 999 se update krdega
   fmt.Println("After updating",slice)
   fmt.Println("Array after update",array)//array bhe change hoga keuki Slice uska reference hai

   //Array ke specific portion ko print krna
   //maan lo hm array ke index 2 se 5 tk print krna chahte hai
   specificSlice := array[2:6]//999,40,50,80

   fmt.Println("specific portion of array",specificSlice)

   //Slice ki length aur capacity check krna
   fmt.Println("Slice Length",len(slice))
   fmt.Println("Slice capacity", cap(slice))

   //ek naya slice bnake usme element copy krna
   newSlice :=make([]int,3)
   copy(newSlice,slice[:3])//pehle 3 element copy krte hai
   fmt.Println("New Slice with Copied Element",newSlice)

   //Slice ke Element ko loop se print krna
   fmt.Println("looping through Slice")
   for i, val := range slice{
	fmt.Printf("index : %d,Value : %d\n",i,val)
   }

   //Element remove krna
   //GO me slice se remove krne ke leyai hm slice ko rearrange krte hai
   indexToRemove :=2
   slice = append(slice[:indexToRemove],slice[indexToRemove+1:]...)
   fmt.Println("After Removing element at index 2:",slice)
}