package main

import "fmt"

// here get salary is used multiple times
type Employee interface {
	CalculateSalary() int
}

type Fulltime struct {
	month  int
	salary int
}

type Contractor struct {
	month  int
	salary int
}

type Freelancer struct {
	hours  int
	salary int
}

func CalculateSalary(employee Employee) int {
	return employee.CalculateSalary()
}
func (f Fulltime) CalculateSalary() int {
	return f.month * 15000
}

func (c Contractor) CalculateSalary() int {
	return c.month * 3000
}

func (g Freelancer) CalculateSalary() int {
	if g.hours > 20 {
		return g.hours * 2000
	}
	return 0
}
func main() {
	var f Fulltime
	var g Freelancer
	var c Contractor
	f.month = 11
	g.hours = 32
	c.month = 13
	fmt.Printf("Fulltime salary is %v\n", CalculateSalary(f))
	fmt.Printf("Contractor salary is %v\n", CalculateSalary(g))
	fmt.Printf("freelancer salary is %v\n", CalculateSalary(c))

}
