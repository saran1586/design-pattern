package main

import "fmt"

type Student struct{
id int
name string

}

func (s Student)Clone() Student {
return s
}


func main(){

 Original Object
	student1 := Student{
		Name: "Saran",
		Age:  21,
		Dept: "CSE",
	}

	
	student2 := student1.Clone()
	student2.Name = "Rahul"
	fmt.Println("Original:", student1)
	fmt.Println("Clone   :", student2)
}




}
