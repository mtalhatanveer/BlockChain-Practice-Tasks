package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollnumber int
	name string
	address string
	subjects []string
}

func NewStudent(rollno int, name string, address string, subjects []string)*Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

type StudentList struct {
	list []*Student
}

func(ls *StudentList) CreateStudent(rollno int, name string, address string, subjects []string)*Student {
	st := NewStudent(rollno, name, address, subjects)
	ls.list = append(ls.list, st)
	return st
}

func CalculateHash (stringToHash string) string{
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}


func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAA", []string{"Maths", "Chemistry", "Physics"})
	student.CreateStudent(25, "Naveed", "BBBBBB", []string{"Maths", "Chemistry", "Physics"})
	fmt.Println(student.list[0].rollnumber)
	
	for i := 0; i <2; i++ {
		var fullBlock string
		fmt.Printf("\n%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student roll no \t %v \n", student.list[i].rollnumber)
		fmt.Printf("Student name \t\t %v \n", student.list[i].name)
		fmt.Printf("Student address \t %v \n", student.list[i].address)
		fmt.Printf("Student subjects \t")
		fullBlock += strconv.Itoa(student.list[i].rollnumber) + student.list[i].name + student.list[i].address
		for j := 0; j<3; j++ {
			fmt.Printf(" %v", student.list[i].subjects[j])
			fullBlock += student.list[i].subjects[j]
		}
		output := CalculateHash(fullBlock)	
		fmt.Printf("\nHash: %x\n", output)
	}
	// fmt.Printf(fullBlock)
}