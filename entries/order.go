package entries

import (
	"fmt"
	"time"
)

type OrderStatus int

const (
	Planned OrderStatus = iota
	NotPlanned
	InProgress
	Done
)

type Order struct {
	jobList         []*Job
	Brigade         *Brigade
	OrderStatus     OrderStatus
	CreationDate    time.Time
	RealizationDate time.Time
	EndDate         time.Time
}

func NewOrder(status OrderStatus) *Order {

	order := &Order{
		jobList:         []*Job{},
		Brigade:         nil,
		OrderStatus:     status,
		CreationDate:    time.Now(),
		RealizationDate: time.Time{},
		EndDate:         time.Time{},
	}

	return order

}

func (o *Order) AddJob(job Job) {
	o.jobList = append(o.jobList, &job)
}

func (o *Order) SetBrigade(brigade Brigade) bool {
	if o.Brigade == nil {
		o.Brigade = &brigade
		return true
	} else {
		panic("Brigade already exists for order")
		return false
	}
}

func (o *Order) StartOrder() {

	o.RealizationDate = time.Now()

	for _, job := range o.jobList {
		job := job
		time.Sleep(time.Duration(job.JobTime) * time.Second)
		fmt.Println(job.Description)
		fmt.Println(job.JobTime)
		fmt.Println(job.IsDone)
		fmt.Println(job.JobType)
		fmt.Println(time.Duration(job.JobTime) * time.Second)
		job.IsDone = true
	}

	o.EndDate = time.Now()

	fmt.Println("All jobs done!")

}
