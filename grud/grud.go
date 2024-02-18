package grud

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DueDate    time.Time `json:"due_date"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const fileName = "tasks.json"

func NewTaskList() TaskList {
	var tasks TaskList
	LoadTasks(&tasks)
	return tasks
}

func HandleChoice(choice int, tasks *TaskList) {
	switch choice {
	case 1:
		AddTask(tasks)
		SaveTasks(tasks)

	case 2:
		FindTasks(tasks)
	case 3:
		UpdateTask(tasks)
		SaveTasks(tasks)

	case 4:
		DeleteTask(tasks)
		SaveTasks(tasks)

	case 5:
		SaveTasks(tasks)
		fmt.Println("Exit")
		os.Exit(0)
	default:
		fmt.Println("Incorrect method")
	}
}

func ShowMenu() {
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Update Task")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}

func LoadTasks(tasks *TaskList) {
	file, err := os.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal(file, tasks)
		if err != nil {
			fmt.Println("Error loading tasks:", err)
		}
	}
}

func SaveTasks(tasks *TaskList) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error saving tasks to file:", err)
	}
}

func AddTask(tasks *TaskList) {
	var title, status string
	var dueDateStr string

	fmt.Print("Enter task title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter task status: ")
	fmt.Scanln(&status)

	fmt.Print("Enter due date (YYYY-MM-DD): ")
	fmt.Scanln(&dueDateStr)
	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Task will be added without a due date.")
		dueDate = time.Time{}
	}

	newTask := Task{
		ID:         len(tasks.Tasks) + 1,
		Title:      title,
		Status:     status,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DueDate:    dueDate,
	}

	tasks.Tasks = append(tasks.Tasks, newTask)
	fmt.Println("Task successfully added.")
}

func FindTasks(tasks *TaskList) {
	if len(tasks.Tasks) == 0 {
		fmt.Println("Task list is empty.")
		return
	}

	var filterStatus string
	fmt.Print("Enter status for filtering (or press Enter to skip): ")
	fmt.Scanln(&filterStatus)

	fmt.Println("Task list:")
	for _, task := range tasks.Tasks {
		if filterStatus == "" || task.Status == filterStatus {
			dueDateStr := ""
			if !task.DueDate.IsZero() {
				dueDateStr = task.DueDate.Format("2006-01-02")
			}

			fmt.Printf("%d. %s (Status: %s, Due Date: %s, Created At: %s, Updated At: %s)\n",
				task.ID, task.Title, task.Status, dueDateStr, task.CreateTime.Format("2006-01-02 15:04:05"), task.UpdateTime.Format("2006-01-02 15:04:05"))
		}
	}
}

func UpdateTask(tasks *TaskList) {
	var id int
	var newStatus string
	var newDueDateStr string

	FindTasks(tasks)

	fmt.Print("Enter task number to update status: ")
	fmt.Scanln(&id)

	if id < 1 || id > len(tasks.Tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	fmt.Print("Enter new status for the task: ")
	fmt.Scanln(&newStatus)

	fmt.Print("Enter new due date (YYYY-MM-DD) or press Enter to skip: ")
	fmt.Scanln(&newDueDateStr)
	newDueDate, err := time.Parse("2006-01-02", newDueDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Due date will not be updated.")
		newDueDate = time.Time{}
	}

	task := &tasks.Tasks[id-1]
	task.Status = newStatus
	task.DueDate = newDueDate
	task.UpdateTime = time.Now()

	fmt.Println("Task status and due date successfully updated.")
}

func DeleteTask(tasks *TaskList) {
	var id int

	FindTasks(tasks)

	fmt.Print("Enter task number to delete: ")
	fmt.Scanln(&id)

	if id < 1 || id > len(tasks.Tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks.Tasks = append(tasks.Tasks[:id-1], tasks.Tasks[id:]...)
	fmt.Println("Task successfully deleted.")
}
