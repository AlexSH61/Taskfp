package main

import (
	"fmt"

	"github.com/AlexSH61/first-project/grud"
)

func main() {
	var tasks grud.TaskList
	grud.LoadTasks(&tasks)

	number := 0
	for {
		grud.ShowMenu()
		fmt.Scanln(&number)
		grud.HandleChoice(number, &tasks)
		if number == 5 {
			break
		}
	}
}
