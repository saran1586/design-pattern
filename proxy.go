package main

import "fmt"

type Student struct{}

func (Student) Marks() {
	fmt.Println("Marks: 95")
}

type Student struct {
	student Student
	login   bool
}

func (s Student) Marks() {
	if s.login {
		s.student.Marks()
	} else {
		fmt.Println("Please Login First")
	}
}

func main() {

	s1 := Student{
		student: Student{},
		login:   true,
	}

	s2 := Student{
		student: Student{},
		login:   false,
	}

	s1.Marks()
	s2.Marks()
}
