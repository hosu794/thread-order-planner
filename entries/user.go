package entries

import (
	"strings"
	"time"
)

type User struct {
	Employee Employee
	login    string
	password string
	initial  string
}

func NewUser(name string, surname string, birthDate time.Time, department WorkerDepartment, login string, password string) *User {
	employee := NewEmployee(name, surname, birthDate, department)

	user := &User{
		*employee,
		login,
		password,
		"",
	}

	user.createInitials()

	return user
}

func (e *User) GetName() string {
	return e.Employee.Name
}

func (e *User) GetSurname() string {
	return e.Employee.Surname
}

func (e *User) GetBirthDate() time.Time {
	return e.Employee.BirthDate
}

func (e *User) GetDepartment() WorkerDepartment {
	return e.Employee.Department
}

func (e *User) GetLogin() string {
	return e.login
}

func (e *User) GetPassword() string {
	return e.password
}

func (e *User) createInitials() {
	firstLetter := strings.ToUpper(e.GetName()[:1])
	secondLetter := strings.ToUpper(e.GetSurname()[:1])

	e.initial = firstLetter + secondLetter
}
