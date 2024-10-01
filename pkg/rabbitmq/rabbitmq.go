package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/imearth/poc-go-dbf/config"
	"github.com/imearth/poc-go-dbf/pkg/manager"
	"github.com/imearth/poc-go-dbf/pkg/services" // Import the services package
	"github.com/streadway/amqp"
)

// Connect to RabbitMQ
func connectRabbitMQ() (*amqp.Connection, error) {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL is not set in .env")
	}

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}

// Publish a message to the queue
func publishMessage(channel *amqp.Channel, queueName string, message string) error {
	err := channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}

// Consume messages from the queue
func consumeMessages(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	messages, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}
	return messages, nil
}

func ConnectRabbit() {
	// Step 1: Connect to RabbitMQ
	conn, err := connectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Step 2: Create a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	expressQueue, expressToERP := config.GetQueueConfigs()
	fmt.Printf("Express Queue: %v\n", expressQueue, expressToERP)
	// Step 3: Declare a queue
	queueName := expressQueue
	// _, err = channel.QueueDeclare(
	// 	queueName, // name
	// 	false,     // durable
	// 	false,     // delete when unused
	// 	false,     // exclusive
	// 	false,     // no-wait
	// 	nil,       // arguments
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to declare a queue: %v", err)
	// }

	// Step 4: Publish a message
	// message := "Hello, RabbitMQ!"
	// err = publishMessage(channel, queueName, message)
	// if err != nil {
	// 	log.Fatalf("Failed to publish a message: %v", err)
	// }
	// fmt.Println("Message Published: ", message)

	// Step 5: Consume messages from the queue
	messages, err := consumeMessages(channel, queueName)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// Step 6: Listen for messages
	go func() {
		expressService := &services.ExpressServiceImpl{}
		jobService := &services.JobServiceImpl{}
		customerService := &services.CustomerServiceImpl{}

		// Initialize the MessageManager
		messageManager := manager.NewMessageManager(expressService, jobService, customerService)

		// Process messages from the queue
		for msg := range messages {
			fmt.Printf("Received a message: %s\n", msg.Body)

			// Process the message
			messageManager.ProcessMessage(msg)
		}
	}()

	// Keep the main function running to listen for messages
	forever := make(chan bool)
	fmt.Println("Waiting for messages...")
	<-forever
}
