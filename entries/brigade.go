package entries

type Brigade struct {
	name         string
	foreman      *Foreman
	employeeList []EmployeeGettersSettersInterface
}

func NewBrigade(name string) *Brigade {

	brigade := &Brigade{
		name,
		nil,
		[]EmployeeGettersSettersInterface{},
	}

	return brigade
}

func (b *Brigade) AddEmployee(employee EmployeeGettersSettersInterface) {
	b.employeeList = append(b.employeeList, employee)
}

func (b *Brigade) AddEmployees(employees []EmployeeGettersSettersInterface) {
	b.employeeList = append(b.employeeList, employees...)
}

func (b *Brigade) SetForeman(foreman Foreman) {
	b.foreman = &foreman
}
