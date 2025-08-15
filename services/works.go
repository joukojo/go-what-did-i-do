// Package services provides various application services.
package services

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joukojo/go-what-did-i-do/fileutil"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

// Work represents a work item in the system.
type Work struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	TaskID      int64      `json:"task_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}

// Works is a collection of Work items.
type Works []Work

// WorksStorage is an in-memory storage for Work items.
var WorkStorage = Works{}

// Add work item to storage
func (ws *Works) Add(work Work) {
	*ws = append(*ws, work)
}

// GetByID retrieves a work item by its ID.
func (ws *Works) GetByID(id int64) *Work {
	for _, work := range *ws {
		if work.ID == id {
			return &work
		}
	}
	return nil
}

func (ws *Works) Exists(id int64) bool {
	return ws.GetByID(id) != nil
}

// LoadWorks loads works from a JSON file.
func (ws *Works) LoadWorks() error {
	data, err := fileutil.GetDataFile("works.json")

	if err != nil {
		return err
	}

	return json.Unmarshal(data, &WorkStorage)
}

func (ws *Works) Print() {

	table := tablewriter.NewTable(os.Stdout, tablewriter.WithConfig(tablewriter.Config{
		Row: tw.CellConfig{
			Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal}, // Wrap long content
			Alignment:    tw.CellAlignment{Global: tw.AlignLeft},     // Left-align rows
			ColMaxWidths: tw.CellWidth{Global: 80},
		},
		Footer: tw.CellConfig{
			Alignment: tw.CellAlignment{Global: tw.AlignRight},
		},
	}),
	)

	table.Header([]string{"ID", "Task", "Description", "Start", "End", "Duration"})
	for _, work := range *ws {

		endDate := "N/A"
		if work.EndDate != nil {
			endDate = work.EndDate.String()
		}

		duration := "N/A"
		if work.EndDate != nil {
			duration = work.EndDate.Sub(work.StartDate).String()
		}
		var taskName = "Unknown"
		if TaskStorage.Exists(work.TaskID) {
			task := TaskStorage.Get(work.TaskID)
			if task != nil {
				taskName = task.Name
			}
		}
		table.Append([]string{
			fmt.Sprintf("%d", work.ID),
			taskName,
			work.Description,
			work.StartDate.String(),
			endDate,
			duration,
		})
	}
	table.Render()
}

func (ws *Works) SaveWorks() error {
	return fileutil.WriteDataFile("works.json", *ws)
}
