package entries

import "time"

var EmployeesList []Employee

type Employee struct {
	Name       string
	Surname    string
	BirthDate  time.Time
	Department WorkerDepartment
}

func NewEmployee(name string, surname string, birthDate time.Time, department WorkerDepartment) (*Employee, error) {

	employee := Employee{
		Name:       name,
		Surname:    surname,
		BirthDate:  birthDate,
		Department: department,
	}

	EmployeesList = append(EmployeesList, employee)

	return &employee, nil
}

type ByFields []Employee

func (a ByFields) Len() int      { return len(a) }
func (a ByFields) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByFields) Less(i, j int) bool {
	if a[i].Name != a[j].Name {
		return a[i].Name < a[j].Name
	}
	if a[i].Surname != a[j].Surname {
		return a[i].Surname < a[j].Surname
	}
	return a[i].BirthDate.Before(a[j].BirthDate)
}
