package services

import (
	"fmt"

	"bytes"
	"context"

	"github.com/imearth/poc-go-dbf/config"
	"github.com/imearth/poc-go-dbf/pkg/manager"
	"github.com/streadway/amqp"

	"log"
	"os/exec"
	"time"
)

// CustomerServiceImpl is the concrete implementation of the CustomerService interface
type ExpressServiceImpl struct{}

// ProcessMessage handles the Customer-related message processing
func (c *ExpressServiceImpl) ProcessReIndex(msg amqp.Delivery) {

	fmt.Println("Processing Customer message in ProcessReIndex...")

	message := manager.ConvertMessageToMessageType(msg)

	fmt.Printf("Processing express message: %s\n", message.Topic)

	runExecutable(
		config.RE_INDEX_EXEC_PATH,
		[]string{config.PROJECT_NAME},
		config.BASE_EXEC_PATH,
		"99",
	)

	msg.Ack(true)

	// Add your customer processing logic here
}

func (c *ExpressServiceImpl) ProcessBackup(msg amqp.Delivery) {
	fmt.Println("Processing Customer message in ProcessBackup...")
	// Add your customer processing logic here
}

func runExecutable(executablePath string, args []string, workingDirectory string, input string) (string, error) {
	// Kill the processes before starting a new one (equivalent to killProcessByName in TypeScript)
	err := killProcessByName("ADM32.EXE")
	if err != nil {
		fmt.Printf("Failed to kill process: %v", err)
	}

	err = killProcessByName("ExpressI.exe")
	if err != nil {
		fmt.Printf("Failed to kill process: %v", err)
	}

	// Prepare the command
	cmd := exec.Command(executablePath, args...)
	cmd.Dir = workingDirectory
	cmd.Stdin = bytes.NewBufferString(input + "\n") // Send input to stdin

	// Set up a channel to handle the process completion
	resultChan := make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Run the command in a goroutine
	go func() {
		err := cmd.Run()
		if err != nil {
			log.Printf("Process finished with error: %v", err)
			resultChan <- "error"
			return
		}

		// Check the exit code of the process
		if cmd.ProcessState.ExitCode() == 0 {
			resultChan <- "completed"
		} else {
			resultChan <- fmt.Sprintf("completed-%d", cmd.ProcessState.ExitCode())
		}
	}()

	// Wait for either process completion or timeout
	select {
	case result := <-resultChan:
		return result, nil
	case <-ctx.Done():
		// Timeout occurred, kill the process
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return "timeout", nil
	}
}

// Dummy function to simulate process killing
func killProcessByName(processName string) error {
	// Here you would implement logic to find and kill the process by name
	// Using OS-specific commands (for example `taskkill` in Windows or `pkill` in Unix-based systems)
	fmt.Printf("Killing process: %s", processName)
	return nil
}
