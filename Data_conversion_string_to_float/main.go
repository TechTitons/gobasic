package main
import ( 
	"fmt"
	"strconv"
)
func main(){
    var text string="79.8"
	
	result,_:=strconv.ParseFloat(text,64)
	
	fmt.Println("Result",result)
}