# State

The "state" design pattern allows an object to change its behavior when its internal state changes. In essence, it appears as if the object changed its class. Instead of having a monolithic object that tries to manage every possible state through conditionals or switch statements, the state pattern promotes the idea of encapsulating state-specific behavior into separate state classes.

Let's look at a basic breakdown:

 - State Interface: Defines an interface to encapsulate the associated behaviors with a particular state of the Context.
 - Concrete States: Classes that implement the State interface and provide the behavior specific to a state.
 - Context: The main object whose behavior changes based on its current state.

 ## Example

Let's consider the scenario of an online order system where an order can be in various states, like New, Shipped, Delivered, or Canceled. Depending on the current state of an order, the operations you can perform on it and the transitions to other states might differ.

```go
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
```

In this example, the `Order` context delegates all its actions to the current state. Each state class handles its respective logic. If a new state or behavior is introduced, we only need to add or modify the corresponding state class, keeping the `Order` context and other states untouched. This makes the system more maintainable and less error-prone.
