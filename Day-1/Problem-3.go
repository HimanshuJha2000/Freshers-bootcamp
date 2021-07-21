package main

import "fmt"

type Employee interface {
	totalSalary() int
}

type PermanentEmployee struct {
	name string
	noOfDays int
	basicPay int
}

func (p PermanentEmployee) getSalary() int {
	return p.basicPay*p.noOfDays
}

type ContractualEmployee struct {
	name string
	noOfDays int
	basicPay int
}

func (c ContractualEmployee) getSalary() int {
	return c.basicPay*c.noOfDays
}

type FreelancerEmployee struct {
	name string
	ratePerHour int
	noOfHours  int
}

func (f FreelancerEmployee) getSalary() int {
	return f.ratePerHour*f.noOfHours
}

func main() {
	permanent := PermanentEmployee{name: "Himanshu", noOfDays: 25, basicPay: 500}
	contractual := ContractualEmployee{name: "Ankit", noOfDays: 30, basicPay: 100}
	freelancer := FreelancerEmployee{name: "Shivam", noOfHours: 44, ratePerHour: 10}

	fmt.Println(permanent.getSalary())
	fmt.Println(contractual.getSalary())
	fmt.Println(freelancer.getSalary())
}





