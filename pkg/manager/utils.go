package manager

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func ConvertMessageToMessageType(msg amqp.Delivery) MessageType {
	var payload MessageType
	err := json.Unmarshal(msg.Body, &payload)
	fmt.Printf("convert message: %s\n", payload)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
	}
	return payload
}
