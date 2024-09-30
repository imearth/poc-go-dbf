package mypackage

import (
	"fmt"
	"log"

	"github.com/tadvi/dbf"
)

// CreateDBF creates a new DBF file and adds records to it
func CreateDBF(filename string) {
	// Create a new DBF table
	table := dbf.New()

	// Add fields to the DBF table
	err := table.AddTextField("NAME", 20) // Text field for name
	if err != nil {
		log.Fatalf("Error adding text field: %v", err)
	}

	err = table.AddNumberField("AGE", 3, 0) // Number field for age
	if err != nil {
		log.Fatalf("Error adding number field: %v", err)
	}

	err = table.AddBoolField("ACTIVE") // Boolean field for active status
	if err != nil {
		log.Fatalf("Error adding boolean field: %v", err)
	}

	// Add the first record
	recordIndex := table.AddRecord()                // Adds a new empty record and returns its index
	table.SetFieldValue(recordIndex, 0, "John Doe") // Set the NAME field value
	table.SetFieldValue(recordIndex, 1, "30")       // Set the AGE field value
	table.SetFieldValue(recordIndex, 2, "T")        // Set the ACTIVE field value

	// Add the second record
	recordIndex = table.AddRecord()                   // Adds another new empty record
	table.SetFieldValue(recordIndex, 0, "Jane Smith") // Set the NAME field value
	table.SetFieldValue(recordIndex, 1, "25")         // Set the AGE field value
	table.SetFieldValue(recordIndex, 2, "F")          // Set the ACTIVE field value

	// Save the DBF file
	err = table.SaveFile(filename)
	if err != nil {
		log.Fatalf("Failed to save DBF file: %v", err)
	}

	fmt.Println("DBF file created successfully")
}

// UpdateDBF updates an existing DBF file if a record's NAME matches the given value
func UpdateDBF(filename string, nameToMatch string, newName string, newAge string) {
	// Load the existing DBF file
	table, err := dbf.LoadFile(filename)
	if err != nil {
		log.Fatalf("Failed to load DBF file: %v", err)
	}

	// Iterate over the records and update the matching one
	for i := 0; i < table.NumRecords(); i++ {
		// Get the value of the NAME field
		name := table.FieldValueByName(i, "NAME")
		if name == nameToMatch {
			// Update the NAME and AGE fields
			table.SetFieldValueByName(i, "NAME", newName)
			table.SetFieldValueByName(i, "AGE", newAge)

			fmt.Printf("Record with name '%s' updated to '%s' with age '%s'\n", nameToMatch, newName, newAge)
		}
	}

	// Save the updated DBF file
	err = table.SaveFile(filename)
	if err != nil {
		log.Fatalf("Failed to save updated DBF file: %v", err)
	}

	fmt.Println("DBF file updated successfully")
}
