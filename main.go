package main

import (
	"log"
	"sort"
	"thread-order-planner/entries"
	"time"
)

func main() {

	emplBlocker := entries.NewEmployeesInUser()
	orderManager := entries.NewOrderManager()

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

	employee4 := entries.NewEmployee("Krzysztof", "Igryski", time.Now(), *department2)

	var employeesForJobOne []entries.EmployeeGettersSettersInterface

	employeesForJobOne = append(employeesForJobOne, employee1)
	employeesForJobOne = append(employeesForJobOne, employee2)
	employeesForJobOne = append(employeesForJobOne, employee3)
	employeesForJobOne = append(employeesForJobOne, expert1)

	var employeesForJobTwo []entries.EmployeeGettersSettersInterface

	employeesForJobTwo = append(employeesForJobTwo, employee4)

	sort.Sort(entries.ByFields(employeesForJobOne))

	foreman1 := entries.NewForeman("Adam", "Kowalksi", time.Now(), *department1, "adam.kowalski", "password")

	foreman2 := entries.NewForeman("Adam", "Nowak", time.Now(), *department1, "adam.nowak", "password")

	brigade1 := entries.NewBrigade("programisci")
	brigade1.SetForeman(*foreman1)

	brigade2 := entries.NewBrigade("testerzy")
	brigade2.SetForeman(*foreman2)

	brigade1.AddEmployees(employeesForJobOne)
	brigade2.AddEmployees(employeesForJobTwo)

	job1 := entries.NewJob(entries.Installation, 5, "opis 123")
	job2 := entries.NewJob(entries.Installation, 3, "opis 123323")

	order1 := entries.NewOrder(entries.Planned)

	job4 := entries.NewJob(entries.Installation, 5, "opis testowe 2")
	job5 := entries.NewJob(entries.Installation, 10, "opisek 232132")

	order2 := entries.NewOrder(entries.Planned)

	order1.AddJob(*job1)
	order1.AddJob(*job2)
	order1.SetBrigade(*brigade1)
	order1.SetBlocker(emplBlocker)

	order2.AddJob(*job4)
	order2.AddJob(*job5)
	order2.SetBrigade(*brigade2)
	order2.SetBlocker(emplBlocker)

	orderManager.AddOrder(order1)
	orderManager.AddOrder(order2)

	orderManager.StartOrdersConcurrently()

}
