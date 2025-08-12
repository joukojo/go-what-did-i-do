// Package services provides various application services.
package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joukojo/go-what-did-i-do/fileutil"
	"github.com/olekukonko/tablewriter"
)

// Task represents a task with a name and description.
type Task struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   int64  `json:"project_id"`
}

// Projects is an array of Project for JSON serialization and deserialization.
// It is used to manage multiple projects in memory.
type Tasks []Task

// TaskStorage is an in-memory store for tasks.
var TaskStorage = Tasks{}

func (ts *Tasks) Add(task Task) {
	*ts = append(*ts, task)
}

// Load tasks from a JSON file.
func (ts *Tasks) LoadTasks() error {
	data, err := fileutil.GetDataFile("tasks.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &TaskStorage)
}

// SaveTasks saves tasks to the data directory.
func (ts *Tasks) SaveTasks() error {
	return fileutil.WriteDataFile("tasks.json", *ts)
}

func (ts *Tasks) PrintTasks() {

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Project name", "Name", "Description"})

	for _, task := range *ts {
		var projectName = "Unknown"
		if project := ProjectStorage.Get(task.ProjectID); project != nil {
			projectName = project.Name
		}
		table.Append([]string{
			fmt.Sprintf("%d", task.ID),
			projectName,
			task.Name,
			task.Description,
		})

	}
	err := table.Render() // Print the table to stdout
	if err != nil {
		fmt.Println("Error rendering table:", err)
	}

}

func (ts *Tasks) Exists(taskID int64) bool {
	for _, task := range *ts {
		if task.ID == taskID {
			return true
		}
	}
	return false
}

// Get retrieves a task by its ID.
func (ts *Tasks) Get(taskID int64) *Task {
	for _, task := range *ts {
		if task.ID == taskID {
			return &task
		}
	}
	return nil
}

// Delete removes a task by its ID.
func (ts *Tasks) Delete(taskID int64) bool {
	for i, task := range *ts {
		if task.ID == taskID {
			*ts = append((*ts)[:i], (*ts)[i+1:]...)
			return true
		}
	}
	return false
}
