package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4}
	fmt.Println("array is",a)
	var b [4]int = [4]int{4,5} /*partial Initialization-agar sari values nahe di to baki element
	                           default value(zero value) laite hai(e.g int ke leyai 0,string ke leyai "")*/
	fmt.Println("array is ",b)

}