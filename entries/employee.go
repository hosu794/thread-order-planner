package entries

import "time"

type Employee struct {
	Name       string
	Surname    string
	BirthDate  time.Time
	Department WorkerDepartment
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
