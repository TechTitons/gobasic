package main

import "fmt"

func main() {
	studentGrade := make(map[string]int)
	studentGrade["prince"] = 95
	studentGrade["goku"] = 78
	studentGrade["titto"] = 65
	studentGrade["pitto"] = 87

	fmt.Println("Marks of Prince", studentGrade["prince"])
	studentGrade["goku"] = 45
	fmt.Println("Marks of goku", studentGrade["goku"])
	fmt.Println("Marks of titto", studentGrade["titto"])
	fmt.Println("Marks of pitto", studentGrade["pitto"]) /*agar aap kise ka value/name
														change kre to uska output me zero(0)
														btaiga to maan lejeyai ke ho skta hai kese
														ka mark such me zero ho to yai kaise btaiga
														jbke vo map me vo person hai. nahe hota to zero
														btata hai mgar such me uska mark zero aur vo
														person ho bhe
														to es problem ko solve krne ke leyai
														app check kro ge ke vo key present hai ke nhe

	jb bhe app map me kise key ko dalte hai findout krne ke leyai ki uska value kya hai to vo 2 cheeje
	return krta hai 1 value return krta hai
					2 ek boolean value return krta hai(jo ke batata hai ke vo exists krta hai ke nahe)
	for example*/
	grades, exists := studentGrade["zeeshan"]
	fmt.Println("Marks of zeeshan", grades)
	fmt.Println("zeeshan exists are not", exists) //agar zeeshan naam ka person hoga to TRUE daiga
	//nahe to FALSE daiga

	Grades,Exists:=studentGrade["sibu"]
	fmt.Println("Marks of sibu",Grades)
	fmt.Println("sibu exists are not",Exists)

	GRades,EXists:=studentGrade["pitto"]
	fmt.Println("Marks of pitto",GRades)
	fmt.Println("pitto exists are not",EXists)

	/*range function aap ka return krta hai index aur value
	lekin map ke case me index key hota hai aur value to value he hota hai
	For Example*/
	for index,value:=range studentGrade{
		fmt.Printf("Key is %s and Marks is %d\n",index,value)
	}

	/*creating a map using a literal(map bnate wakt initialize kr daina)*/
	person :=map[string]int{
	"Alice":90,
	"Bob":85,
	"Charle":87,
	}
	for index,value:=range person{
		fmt.Printf("--------------Key is %s and Marks is %d\n",index,value)
	}


}
