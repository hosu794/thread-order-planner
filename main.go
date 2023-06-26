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

	employee1, err := entries.NewEmployee("Grzegorz", "Szczesny", time.Now(), *department1)

	if err != nil {
		log.Println(employee1)
		panic(err)
	}

	employee2, err := entries.NewEmployee("Agata", "Bąk", time.Now(), *department2)

	if err != nil {
		log.Println(employee2)
		panic(err)
	}

	employee3, err := entries.NewEmployee("Krzysztof", "Kowalski", time.Now(), *department1)

	if err != nil {
		log.Println(employee3)
		panic(err)
	}

	var employees []entries.Employee

	employees = append(employees, *employee1)
	employees = append(employees, *employee2)
	employees = append(employees, *employee3)

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
