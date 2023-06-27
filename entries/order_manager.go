package entries

import "sync"

type OrderManager struct {
	orders []*Order
}

func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders: []*Order{},
	}
}

func (om *OrderManager) AddOrder(order *Order) {
	om.orders = append(om.orders, order)
}

func (om *OrderManager) StartOrdersConcurrently() {
	var wg sync.WaitGroup

	wg.Add(len(om.orders))

	completionChannel := make(chan struct{})

	for _, order := range om.orders {
		order := order

		go func() {
			order.StartOrder()

			completionChannel <- struct{}{}
		}()
	}

	go func() {
		wg.Wait()

		close(completionChannel)
	}()

	for range completionChannel {
		// Decrement the wait group counter
		wg.Done()
	}

}
