package entries

type EmployeeBlocker struct {
	employees []EmployeeGettersSettersInterface
}

func NewEmployeesInUser() *EmployeeBlocker {
	return &EmployeeBlocker{
		employees: []EmployeeGettersSettersInterface{},
	}
}

func (e *EmployeeBlocker) IsContains(elements []EmployeeGettersSettersInterface) bool {

	if len(e.employees) == 0 {
		return false
	}

	for _, blockedEmployee := range e.employees {
		for _, jobEmployee := range elements {
			if blockedEmployee == jobEmployee {
				return true
			}
		}
	}
	return false
}

func (e *EmployeeBlocker) GetEmployees() []EmployeeGettersSettersInterface {
	return e.employees
}

func (e *EmployeeBlocker) AddEmployees(elements []EmployeeGettersSettersInterface) {
	e.employees = append(e.employees, elements...)
}

func (e *EmployeeBlocker) RemoveEmployees(elements []EmployeeGettersSettersInterface) {

	deleteResults := make([]EmployeeGettersSettersInterface, 0)

	elementsMap := make(map[EmployeeGettersSettersInterface]bool)

	for _, el := range elements {
		elementsMap[el] = true
	}

	for _, el := range e.employees {
		if !elementsMap[el] {
			deleteResults = append(deleteResults, el)
		}
	}

	e.employees = deleteResults

}
