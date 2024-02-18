package main

import (
	"fmt"

	grud "github.com/AlexSH61/Taskfp/grud"
)

func main() {
	// Create a TaskList instance
	tasks := grud.NewTaskList()

	for {
		// Display menu
		grud.ShowMenu()

		// Get user choice
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// Handle user choice
		grud.HandleChoice(choice, &tasks)
	}
}
