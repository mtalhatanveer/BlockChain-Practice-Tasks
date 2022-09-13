package main

import "fmt"

type Doctor struct {
	number int
	doctorName string
	patients []string
}

func passStruct(class Doctor){
	fmt.Printf("Doctor ID=%v, Doctor name=%v, Patients=%v", class.number, class.doctorName, class.patients)
}

func main() {
	aDoctor := Doctor{ 
		number: 006,
		doctorName: "Naveed",
		patients: []string{"A", "B", "C", "D"},
	}

	passStruct(aDoctor)
}