// Package main provides the entry point for the 'What did I do?' application.
package main

import (
	"github.com/joukojo/go-what-did-i-do/cmd"
	"github.com/joukojo/go-what-did-i-do/fileutil"
	"github.com/joukojo/go-what-did-i-do/services"
)

func initialize() {

	err := fileutil.GetDataDirectory()
	if err != nil {
		panic(err)
	}

	err = services.CustomerStorage.Load()
	if err != nil {
		panic(err)
	}

	err = services.ProjectStorage.LoadProjects()
	if err != nil {
		panic(err)
	}

	err = services.TaskStorage.LoadTasks()
	if err != nil {
		panic(err)
	}

}

func main() {

	initialize()
	cmd.Execute()

}
