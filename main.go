package main

import (
	"fmt"
	"log"
	"sort"
	"thread-order-planner/entries"
	"time"
)

func main() {

	fmt.Println("test")

	department1, err := entries.NewWorkerDepartment("programisci")

	if err != nil {
		log.Println(department1)
		panic(err)
	}

	department2, err := entries.NewWorkerDepartment("hr")

	if err != nil {
		log.Println(department2)
		panic(err)
	}

	employee1 := entries.Employee{
		Name:       "Grzegorz",
		Surname:    "Szczęsny",
		BirthDate:  time.Now(),
		Department: *department1,
	}

	employee2 := entries.Employee{
		Name:       "Agata",
		Surname:    "Bąk",
		BirthDate:  time.Now(),
		Department: *department2,
	}

	employee3 := entries.Employee{
		Name:       "Krzysztof",
		Surname:    "Kowalski",
		BirthDate:  time.Now(),
		Department: *department1,
	}

	var employees []entries.Employee

	employees = append(employees, employee1)
	employees = append(employees, employee2)
	employees = append(employees, employee3)

	fmt.Println("Before sort:")

	for _, employee := range employees {
		fmt.Println(employee.Name, employee.Surname, employee.BirthDate, employee.Department.Name)
	}

	fmt.Println("After sorting")

	sort.Sort(entries.ByFields(employees))

	for _, employee := range employees {
		fmt.Println(employee.Name, employee.Surname, employee.BirthDate, employee.Department.Name)
	}

}
