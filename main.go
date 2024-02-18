package main

import (
	"fmt"

	grud "github.com/AlexSH61/Taskfp/grud"
)

func main() {
	tasks := grud.NewTaskList()

	for {
		grud.ShowMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		grud.HandleChoice(choice, &tasks)
	}
}
