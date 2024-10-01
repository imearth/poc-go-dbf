package services

import (
	"fmt"

	"github.com/streadway/amqp"
)

// CustomerServiceImpl is the concrete implementation of the CustomerService interface
type CustomerServiceImpl struct{}

// ProcessMessage handles the Customer-related message processing
func (c *CustomerServiceImpl) ProcessMessage(msg amqp.Delivery) {
	fmt.Println("Processing message in CustomerService...")

	// Add your customer processing logic here
	msg.Ack(true)
}
