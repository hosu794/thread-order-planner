package entries

import "time"

type Foreman struct {
	User   User
	Orders []Order
}

func NewForeman(name string, surname string, birthDate time.Time, department WorkerDepartment, login string, password string) *Foreman {
	user := NewUser(name, surname, birthDate, department, login, password)

	foreman := &Foreman{
		*user,
		[]Order{},
	}

	return foreman
}

func (e *Foreman) GetName() string {
	return e.User.GetName()
}

func (e *Foreman) GetSurname() string {
	return e.User.GetName()
}

func (e *Foreman) GetBirthDate() time.Time {
	return e.User.GetBirthDate()
}

func (e *Foreman) GetDepartment() WorkerDepartment {
	return e.User.GetDepartment()
}

func (e *Foreman) GetLogin() string {
	return e.User.GetLogin()
}

func (e *Foreman) GetPassword() string {
	return e.User.GetPassword()
}
