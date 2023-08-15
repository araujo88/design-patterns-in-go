package main

import "fmt"

// State interface
type OrderState interface {
	ShipOrder()
	DeliverOrder()
	CancelOrder()
}

// Concrete States
type NewOrderState struct{}

func (n *NewOrderState) ShipOrder() {
	fmt.Println("Shipping Order.")
}

func (n *NewOrderState) DeliverOrder() {
	fmt.Println("Can't deliver an order which hasn't been shipped!")
}

func (n *NewOrderState) CancelOrder() {
	fmt.Println("Canceling Order.")
}

type ShippedOrderState struct{}

func (s *ShippedOrderState) ShipOrder() {
	fmt.Println("Can't ship an order again.")
}

func (s *ShippedOrderState) DeliverOrder() {
	fmt.Println("Delivering Order.")
}

func (s *ShippedOrderState) CancelOrder() {
	fmt.Println("Can't cancel a shipped order!")
}

type DeliveredOrderState struct{}

func (d *DeliveredOrderState) ShipOrder() {
	fmt.Println("Order already delivered.")
}

func (d *DeliveredOrderState) DeliverOrder() {
	fmt.Println("Order already delivered.")
}

func (d *DeliveredOrderState) CancelOrder() {
	fmt.Println("Can't cancel a delivered order!")
}

// Context
type Order struct {
	state OrderState
}

func (o *Order) ShipOrder() {
	o.state.ShipOrder()
}

func (o *Order) DeliverOrder() {
	o.state.DeliverOrder()
}

func (o *Order) CancelOrder() {
	o.state.CancelOrder()
}

func NewOrder() *Order {
	return &Order{state: &NewOrderState{}}
}

func main() {
	order := NewOrder()
	order.ShipOrder()
	order.DeliverOrder()

	// Setting a new state manually just for demonstration.
	order.state = &ShippedOrderState{}
	order.CancelOrder()
}
