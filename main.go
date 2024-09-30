// main.go
package main

import (
	"fmt"

	"github.com/imearth/poc-go-dbf/pkg/mypackage"
)

func main() {
	fmt.Println("Hello from main!")
	mypackage.MyFunction()
	mypackage.CreateDBF("test.dbf")
	mypackage.UpdateDBF("test.dbf", "John Doe", "John Doe Updated", "31")

}
