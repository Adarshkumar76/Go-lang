package main

import (
	"fmt"
	"time"
)

// Order data struct
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func NewOrder(id string, amount float32, status string) *Order {
	return &Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// changeStatus method to update the status of an Order
func (o *Order) changeStatus(status string) {
	o.status = status
}

func main() {
	// Example of creating an Order instance
	order1 := NewOrder("12345", 99, "Pending")
	fmt.Println("Order ID:", order1.id)
	fmt.Println("Amount:", order1.amount)
	fmt.Println("Status:", order1.status)
	myOrder := Order{
		id:        "12345",
		amount:    99.99,
		status:    "Pending",
		createdAt: time.Now(),
	}
	myOrder.changeStatus("Shipped")
	fmt.Println("Order ID:", myOrder)

	// Note: In a real application, you would typically use a constructor function or similar pattern
	// to create instances of structs.
	// Here, we are directly initializing the struct for demonstration purposes.

	// var order Order
	// order.id = "12345"
	// order.amount = 99.99
	// order.status = "Pending"
	// order.createdAt = time.Now()

	// // Example usage of the Order struct
	// println("Order ID:", order.id)
	// println("Amount:", order.amount)
	// println("Status:", order.status)
	// println("Created At:", order.createdAt.String())

}
