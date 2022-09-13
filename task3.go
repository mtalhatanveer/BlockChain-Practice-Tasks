package main

import (
	"fmt"
	"strings"
)


type Student struct {
	rollnumber int
	name string
	address string
}

func NewStudent(rollno int, name string, address string)*Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func(ls *StudentList) CreateStudent(rollno int, name string, address string)*Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")
	fmt.Println(student.list[0].rollnumber)

	for i := 0; i <2; i++ {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student roll no \t %v \n", student.list[i].rollnumber)
		fmt.Printf("Student name \t\t %v \n", student.list[i].name)
		fmt.Printf("Student address \t %v \n", student.list[i].address)
	}
}