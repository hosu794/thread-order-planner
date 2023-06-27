package entries

import "time"

type EmployeeGettersSettersInterface interface {
	GetName() string
	GetSurname() string
	GetBirthDate() time.Time
	GetDepartment() WorkerDepartment
}

var EmployeesList []Employee

type Employee struct {
	Name       string
	Surname    string
	BirthDate  time.Time
	Department WorkerDepartment
}

func NewEmployee(name string, surname string, birthDate time.Time, department WorkerDepartment) *Employee {

	employee := Employee{
		Name:       name,
		Surname:    surname,
		BirthDate:  birthDate,
		Department: department,
	}

	EmployeesList = append(EmployeesList, employee)

	return &employee
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetSurname() string {
	return e.Surname
}

func (e *Employee) GetBirthDate() time.Time {
	return e.BirthDate
}

func (e *Employee) GetDepartment() WorkerDepartment {
	return e.Department
}

type ByFields []EmployeeGettersSettersInterface

func (a ByFields) Len() int      { return len(a) }
func (a ByFields) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByFields) Less(i, j int) bool {
	if a[i].GetName() != a[j].GetName() {
		return a[i].GetName() < a[j].GetName()
	}
	if a[i].GetSurname() != a[j].GetSurname() {
		return a[i].GetSurname() < a[j].GetSurname()
	}
	return a[i].GetBirthDate().Before(a[j].GetBirthDate())
}
