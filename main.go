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

	employee1 := entries.NewEmployee("Grzegorz", "Szczesny", time.Now(), *department1)

	employee2 := entries.NewEmployee("Agata", "Bąk", time.Now(), *department2)

	employee3 := entries.NewEmployee("Krzysztof", "Kowalski", time.Now(), *department1)

	expert1 := entries.NewExpert("Adam", "Kozuch", time.Now(), *department1, "Senior Golang Developer")

	var employees []entries.EmployeeGettersSettersInterface

	employees = append(employees, employee1)
	employees = append(employees, employee2)
	employees = append(employees, employee3)
	employees = append(employees, expert1)

	sort.Sort(entries.ByFields(employees))

	foreman1 := entries.NewForeman("Adam", "Kowalksi", time.Now(), *department1, "adam.kowalski", "password")

	brigade1 := entries.NewBrigade("programisci")
	brigade1.SetForeman(*foreman1)

	brigade1.AddEmployees(employees)

	job1 := entries.NewJob(entries.Installation, 5, "opis 123")
	job2 := entries.NewJob(entries.Installation, 3, "opis 123323")

	order1 := entries.NewOrder(entries.Planned)

	order1.AddJob(*job1)
	order1.AddJob(*job2)

	order1.SetBrigade(*brigade1)

	order1.StartOrder()

}
