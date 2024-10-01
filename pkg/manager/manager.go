package manager

import (
	"fmt"

	"github.com/streadway/amqp"
)

// NewMessageManager is the constructor for MessageManager
func NewMessageManager(expressService ExpressService, jobService JobService, customerService CustomerService) *MessageManager {
	return &MessageManager{
		expressService:  expressService,
		jobService:      jobService,
		customerService: customerService,
	}
}

func (m *MessageManager) ProcessMessage(msg amqp.Delivery) {

	fmt.Printf("Processing message: %s\n", msg.Body)

	message := ConvertMessageToMessageType(msg)

	switch message.Topic {
	case "Job":
		m.jobService.ProcessMessage(msg)
	case "Customer":
		m.customerService.ProcessMessage(msg)
	case "ReIndex":
		m.expressService.ProcessReIndex(msg)
	case "Backup":
		m.expressService.ProcessBackup(msg)
	default:
		fmt.Printf("Unknown topic: %s\n", message.Topic)
	}
}
