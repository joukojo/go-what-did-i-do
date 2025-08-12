// Package services provides functionality for managing customers.
package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joukojo/go-what-did-i-do/fileutil"

	"github.com/olekukonko/tablewriter"
)

// Customer structure for json serialization and deserialization.
type Customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Customers array for JSON serialization and deserialization.
type Customers []Customer

// CustomerStorage in memory
var CustomerStorage Customers

const customerFileName = "customers.json"

// Add customer
func (c *Customers) Add(customer Customer) {
	*c = append(*c, customer)
}

// Get returns a customer by ID.
func (c *Customers) Get(id int64) *Customer {
	for _, customer := range *c {
		if customer.ID == id {
			return &customer
		}
	}
	return nil
}

// Exists checks if a customer exists by ID.
func (c *Customers) Exists(id int64) bool {
	return c.Get(id) != nil
}

// Update customer name by id
func (c *Customers) Update(id int64, name string) bool {
	for i, customer := range *c {
		if customer.ID == id {
			(*c)[i].Name = name
			return true
		}
	}
	return false
}

// Delete customer by id
func (c *Customers) Delete(id int64) bool {
	for i, customer := range *c {
		if customer.ID == id {
			*c = append((*c)[:i], (*c)[i+1:]...)
			return true
		}
	}
	return false
}

// Print prints the customers in a table format.
func (c *Customers) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Name"})
	// table.SetAutoFormatHeaders(true)
	for _, customer := range *c {

		err := table.Append([]string{fmt.Sprintf("%d", customer.ID), customer.Name})
		if err != nil {
			fmt.Println("Error appending to table:", err)
			return
		}
	}
	err := table.Render() // Print the table to stdout
	if err != nil {
		fmt.Println("Error rendering table:", err)
	}

}

// Save customer data to JSON file.
func (c *Customers) Save() error {

	err := fileutil.WriteDataFile(customerFileName, CustomerStorage)
	if err != nil {
		return err
	}
	return nil
}

// Load customers from a JSON file
func (c *Customers) Load() error {
	content, err := fileutil.GetDataFile(customerFileName)

	if err != nil {
		fmt.Println("Error loading customers:", err)
		return err
	}

	customers := Customers{}
	err = json.Unmarshal(content, &customers)
	if err != nil {
		fmt.Println("Error unmarshalling customers:", err)
		return err
	}

	for _, customer := range customers {
		CustomerStorage.Add(customer)
	}

	return nil
}
