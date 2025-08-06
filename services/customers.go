package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joukojo/go-what-did-i-do/fileutil"

	"github.com/olekukonko/tablewriter"
)

type Customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Customers []Customer

var CustomerStorage Customers

const customerFileName = "customers.json"

func (c *Customers) Add(customer Customer) {
	*c = append(*c, customer)
}

/**
 * Get returns a customer by ID.
 */
func (c *Customers) Get(id int64) *Customer {
	for _, customer := range *c {
		if customer.ID == id {
			return &customer
		}
	}
	return nil
}
func (c *Customers) Update(id int64, name string) bool {
	for i, customer := range *c {
		if customer.ID == id {
			(*c)[i].Name = name
			return true
		}
	}
	return false
}
func (c *Customers) Delete(id int64) bool {
	for i, customer := range *c {
		if customer.ID == id {
			*c = append((*c)[:i], (*c)[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Customers) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Name"})
	// table.SetAutoFormatHeaders(true)
	for _, customer := range *c {

		table.Append([]string{fmt.Sprintf("%d", customer.ID), customer.Name})
	}
	table.Render() // Print the table to stdout}
}

func (c *Customers) Save() error {

	err := fileutil.WriteFile(customerFileName, CustomerStorage)
	if err != nil {
		return err
	}
	return nil
}
func (c *Customers) Load() {
	content, err := fileutil.GetDataFile(customerFileName)

	if err != nil {
		fmt.Println("Error loading customers:", err)
		return
	}

	customers := Customers{}
	err = json.Unmarshal(content, &customers)

	for _, customer := range customers {
		CustomerStorage.Add(customer)
	}

}
