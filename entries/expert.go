package entries

import (
	"time"
)

type Expert struct {
	Employee       Employee
	specialization string
}

func NewExpert(name string, surname string, birthDate time.Time, department WorkerDepartment, specialization string) *Expert {
	employee := NewEmployee(name, surname, birthDate, department)

	expert := &Expert{
		*employee,
		specialization,
	}

	return expert
}

func (e *Expert) GetName() string {
	return e.Employee.Name
}

func (e *Expert) GetSurname() string {
	return e.Employee.Surname
}

func (e *Expert) GetBirthDate() time.Time {
	return e.Employee.BirthDate
}

func (e *Expert) GetDepartment() WorkerDepartment {
	return e.Employee.Department
}
