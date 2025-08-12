// Package services provides functionality for managing projects.
package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joukojo/go-what-did-i-do/fileutil"
	"github.com/olekukonko/tablewriter"
)

// Project represents a project with a name and description.
type Project struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CustomerID  int64  `json:"customer_id"`
}

// Projects is an array of Project for JSON serialization and deserialization.
// It is used to manage multiple projects in memory.
type Projects []Project

// ProjectStorage is an in-memory store for projects.
var ProjectStorage = Projects{}

// Add project
func (ps *Projects) Add(project Project) {
	*ps = append(*ps, project)
}

// GetDataDirectory retrieves the data directory for the application.
func (ps *Projects) GetDataDirectory() error {
	return fileutil.GetDataDirectory()
}

// LoadProjects loads projects from the data directory.
func (ps *Projects) LoadProjects() error {
	data, err := fileutil.GetDataFile("projects.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &ProjectStorage)
}

// SaveProjects saves projects to the data directory.
func (ps *Projects) SaveProjects() error {
	return fileutil.WriteDataFile("projects.json", *ps)
}

// PrintProjects prints all projects in a table format.
func (ps *Projects) PrintProjects() {

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Customer", "Name", "Description"})
	// table.SetAutoFormatHeaders(true)
	for _, project := range *ps {

		var customerName = "Unknown"
		if project.CustomerID != 0 {
			customer := CustomerStorage.Get(project.CustomerID)
			if customer != nil {
				customerName = customer.Name
			}
		}
		err := table.Append([]string{fmt.Sprintf("%d", project.ID), customerName, project.Name, project.Description})
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

// Get returns a project by ID.
func (ps *Projects) Get(id int64) *Project {
	for _, project := range *ps {
		if project.ID == id {
			return &project
		}
	}
	return nil
}

// Exists checks if a project exists by ID.
func (ps *Projects) Exists(id int64) bool {
	return ps.Get(id) != nil
}

// Delete removes a project by ID.
func (ps *Projects) Delete(id int64) {
	for i, project := range *ps {
		if project.ID == id {
			*ps = append((*ps)[:i], (*ps)[i+1:]...)
			return
		}
	}
}
