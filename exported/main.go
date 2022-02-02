package main

import (
	"exported/staff"
	"log"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, Fulltime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, Fulltime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, Fulltime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, Fulltime: false},
	{FirstName: "Francesca", LastName: "Vania", Salary: 100000, Fulltime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}
	// myStaff.All()

	// log.Println(myStaff.All())
	// Export var must be first uppercase
	// staff.OverPaidLimit = 70000
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("Underpaid staff:", myStaff.Underpaid())
}
