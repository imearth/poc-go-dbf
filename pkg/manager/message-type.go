package manager

import (
	"time"

	"github.com/streadway/amqp"
)

type MessageType struct {
	ID         string    `json:"id"`
	Topic      string    `json:"topic"`
	ActionType *string   `json:"actionType"` // Nullable string
	Payload    *string   `json:"payload"`    // Nullable string
	Status     string    `json:"status"`
	Reason     *string   `json:"reason"` // Nullable string
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ExpressService interface {
	ProcessReIndex(msg amqp.Delivery)
	ProcessBackup(msg amqp.Delivery)
}

type JobService interface {
	ProcessMessage(msg amqp.Delivery)
}

type CustomerService interface {
	ProcessMessage(msg amqp.Delivery)
}

// MessageManager struct to manage the services
type MessageManager struct {
	expressService  ExpressService
	jobService      JobService
	customerService CustomerService
}
