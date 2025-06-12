package main
import "fmt"
func main(){
	num:=[]int{1,2,3,4,5,6,7,8}
	max:=num[0]
	for i:=1;i<len(num);i++{
		if num[i]>max{
			max=num[i]
		}
	}
	fmt.Println("largest",max)
}