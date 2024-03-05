package main

import (
	"fmt"
	"sync"
)
type orderStatus struct {
	id       string
	customer string
	tracking string
	status   string
}

type Warehouse struct {
	Orders map[string]orderStatus
	mu     sync.Mutex
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		Orders: make(map[string]orderStatus),
	}
}

func (w *Warehouse) ProcessOrder(order orderStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	w.mu.Lock()
	w.Orders[order.tracking] = order
	w.mu.Unlock()
}

// logic for updating order status
func (w *Warehouse) UpdateOrderStatus(trackingNumber string, status string, wg *sync.WaitGroup) {
	defer wg.Done()
	w.mu.Lock()
	if order, ok := w.Orders[trackingNumber]; ok {
		order.status = status
		w.Orders[trackingNumber] = order
	}
	w.mu.Unlock()
}

type notif struct{}

func (ns *notif) NotifyCustomer(order orderStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order with ID %s has been Processed\nCustomer %s has been notified about the shipment status: %s with tracking number %s.\n", order.id, order.customer, order.status, order.tracking)
}

func main() {
	warehouse := NewWarehouse()
	var wg sync.WaitGroup
	var notifyWg sync.WaitGroup // agar bisa muncul tiga2nya
	shipments := []orderStatus{
		{id: "1", customer: "Alice", status: "In Transit", tracking: "123-132"},
		{id: "2", customer: "Bob", status: "Packing", tracking: "932-652"},
		{id: "3", customer: "Charlie", status: "Delivered", tracking: "578-852"},
	}
	wg.Add(len(shipments))
	notifyWg.Add(len(shipments))

	for _, shipment := range shipments {
		go func(shipment orderStatus) {
			warehouse.ProcessOrder(shipment, &wg)
			ns := &notif{}
			ns.NotifyCustomer(shipment, &notifyWg) 
		}(shipment)
	}

	wg.Wait()
	notifyWg.Wait() 
}
