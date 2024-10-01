package services

import (
	"fmt"
	"log"

	"github.com/imearth/poc-go-dbf/pkg/manager"
	"github.com/streadway/amqp"
)

// CustomerServiceImpl is the concrete implementation of the CustomerService interface
type JobServiceImpl struct{}

// ProcessMessage handles the Customer-related message processing
func (c *JobServiceImpl) ProcessMessage(msg amqp.Delivery) {
	// Process the message (this is where your actual business logic would go)

	message := manager.ConvertMessageToMessageType(msg)
	fmt.Printf("Processing Job message: %s\n", message.Topic)

	// Acknowledge the message after processing
	err := msg.Ack(false)
	if err != nil {
		log.Printf("Failed to acknowledge message: %v", err)
		return
	}

	fmt.Println("Job message acknowledged.")
}
