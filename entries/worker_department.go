package entries

import "thread-order-planner/errors"

type WorkerDepartment struct {
	Name string
}

var nameMap = make(map[string]bool)

func NewWorkerDepartment(name string) (*WorkerDepartment, error) {

	if exists(name) {
		return nil, errors.NewNotUniqueNameError(name)
	}

	wd := &WorkerDepartment{
		Name: name,
	}

	nameMap[name] = true

	return wd, nil

}

func exists(name string) bool {
	return nameMap[name]
}
