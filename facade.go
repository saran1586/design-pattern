package main

import "fmt"

type Attendance struct{}

func (Attendance) Check() {
	fmt.Println("attendance")
}

type Marks struct{}

func (Marks) Show() {
	fmt.Println("marks displayed")
}


type Student struct {
	attendance Attendance
	marks      Marks
	
}

func (s Student) StudentDetails() {
	s.attendance.Check()
	s.marks.Show()
	fmt.Println("Student Details Retrieved Successfully")
}

func main() {
	student := Student{}
	student.StudentDetails()
}
