package main
import "fmt"
type Student struct{
	Name string
	Marks float64
}

func printStudent(s Student){
	fmt.Println("Name",s.Name,"Marks",s.Marks)

}

func main(){
	s :=Student{Name: "zeeshan",Marks: 88.8}
	printStudent(s)
}