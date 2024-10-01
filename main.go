// main.go
package main

import (
	"fmt"
	"log"
	"os"

	//"github.com/imearth/poc-go-dbf/pkg/dbf"
	"github.com/imearth/poc-go-dbf/pkg/mypackage"
	"github.com/imearth/poc-go-dbf/pkg/rabbitmq"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file at the start of the application
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	fmt.Println("Hello from main!")

	fmt.Println("RABBITMQ_URL:", os.Getenv("RABBITMQ_URL"))

	mypackage.MyFunction()

	// dbf.CreateDBF("test.dbf")
	// dbf.UpdateDBF("test.dbf", "John Doe", "John Doe Updated", "31")

	rabbitmq.ConnectRabbit()

}
