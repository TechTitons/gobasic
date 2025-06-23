package main
import "fmt"

type Student struct{
	Name string
	Rollno int
	Marks  float64
}
func main(){
	var s Student
	s.Name="Zeeshan"
	s.Rollno=104
	s.Marks=85.4
	fmt.Println("Name",s.Name)
	fmt.Println("Rollno",s.Rollno)
	fmt.Println("Marks",s.Marks)

}