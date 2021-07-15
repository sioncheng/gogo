package main

import "fmt"

type Employee struct {
	Name string
	Id   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.Id)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) Description() string {
	return fmt.Sprintf("manager %s (%s)", m.Name, m.Id)
}

func main() {
	p := Employee{"E", "e1"}
	fmt.Println(p.Description())

	p2 := Employee{"D", "d1"}

	reports := []Employee{p, p2}
	m := Manager{
		Employee: Employee{"M", "m1"},
		Reports:  reports,
	}

	fmt.Println(m)
}
