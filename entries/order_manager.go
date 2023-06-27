package entries

import (
	"fmt"
	"sync"
	"time"
)

type OrderManager struct {
	orders  []*Order
	blocker *EmployeeBlocker
}

func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders:  []*Order{},
		blocker: nil,
	}
}

func (om *OrderManager) AddOrder(order *Order) {
	om.orders = append(om.orders, order)
}

func (om *OrderManager) StartOrdersConcurrently() {

	if om.blocker == nil {
		panic("Blocker not set")
	}

	var wg sync.WaitGroup

	wg.Add(len(om.orders))

	completionChannel := make(chan struct{})

	for _, order := range om.orders {
		order := order

		go func() {

			om.blockOrderIfEmployeesLocked(order)

			om.blocker.AddEmployees(order.Brigade.employeeList)

			order.StartOrder()

			om.blocker.RemoveEmployees(order.Brigade.employeeList)

			completionChannel <- struct{}{}
		}()
	}

	go func() {
		wg.Wait()

		close(completionChannel)
	}()

	for range completionChannel {
		wg.Done()
	}
}

func (om *OrderManager) blockOrderIfEmployeesLocked(order *Order) {
	if om.blocker.IsContains(order.Brigade.employeeList) {
		fmt.Println("Employees are locked")
		for {
			time.Sleep(1 * time.Second)
			if !om.blocker.IsContains(order.Brigade.employeeList) {
				fmt.Println("Employees are free")
				break
			}
		}
	}
}

func (om *OrderManager) SetBlocker(blocker *EmployeeBlocker) {
	om.blocker = blocker
}
