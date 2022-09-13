package main

import "fmt"

type employee struct {

	name string
	
	salary int
	
	position string
	
}
	
type company struct {

	companyName string

	employees []employee

}

func main() {
	emp1 := employee{
		name: "Muhammad",
		salary: 1000,
		position: "CEO",
	}
	emp2 := employee{
		name: "Talha",
		salary: 1000,
		position: "CTO",
	}
	emp3 := employee{
		name: "Tanveer",
		salary: 1000,
		position: "CFO",
	}

	aCompany := company{
		companyName: "Company 123",
		employees: []employee{emp1, emp2, emp3},
	}

	fmt.Printf("Company Name=%v, Employees=%v", aCompany.companyName, aCompany.employees)
	
}