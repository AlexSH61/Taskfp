package grud

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	tasks := NewTaskList()

	// Mock user input
	mockInput := "Test Task\nInProgress\n2024-02-28\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(mockInput))
	r.Close()

	// Call the AddTask function
	AddTask(&tasks)

	// Restore stdin
	os.Stdin = os.NewFile(0, "/dev/stdin")

	// Check if the task was added
	if len(tasks.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks.Tasks))
	}

	// Check task details
	task := tasks.Tasks[0]
	if task.Title != "Test Task" || task.Status != "InProgress" || task.DueDate.Format("2006-01-02") != "2024-02-28" {
		t.Errorf("Unexpected task details: %+v", task)
	}
}

func TestUpdateTask(t *testing.T) {
	tasks := NewTaskList()

	// Add a task for testing
	tasks.Tasks = append(tasks.Tasks, Task{
		ID:         1,
		Title:      "Test Task",
		Status:     "ToDo",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DueDate:    time.Now().AddDate(0, 0, 7),
	})

	// Mock user input
	mockInput := "1\nInProgress\n2024-03-10\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(mockInput))
	r.Close()

	// Call the UpdateTask function
	UpdateTask(&tasks)

	// Restore stdin
	os.Stdin = os.NewFile(0, "/dev/stdin")

	// Check if the task was updated
	if tasks.Tasks[0].Status != "InProgress" || tasks.Tasks[0].DueDate.Format("2006-01-02") != "2024-03-10" {
		t.Errorf("Unexpected task details after update: %+v", tasks.Tasks[0])
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.TempFile("", "tasks_test_*.json")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(tmpFile.Name())

	// Create a new TaskList and add a task
	tasks := NewTaskList()
	tasks.Tasks = append(tasks.Tasks, Task{
		ID:         1,
		Title:      "Test Task",
		Status:     "ToDo",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DueDate:    time.Now().AddDate(0, 0, 7),
	})

	// Save tasks to the temporary file
	SaveTasks(&tasks, tmpFile.Name())

	// Load tasks from the temporary file
	loadedTasks := NewTaskList()
	LoadTasks(&loadedTasks, tmpFile.Name())

	// Check if the loaded tasks match the original tasks
	if len(loadedTasks.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(loadedTasks.Tasks))
	}

	// Check task details
	originalTask := tasks.Tasks[0]
	loadedTask := loadedTasks.Tasks[0]
	if originalTask.Title != loadedTask.Title || originalTask.Status != loadedTask.Status ||
		originalTask.DueDate != loadedTask.DueDate {
		t.Errorf("Loaded task details do not match original: %+v vs %+v", originalTask, loadedTask)
	}
}

// Helper function to create a TaskList with sample tasks
func createSampleTaskList() TaskList {
	return TaskList{
		Tasks: []Task{
			{ID: 1, Title: "Task 1", Status: "ToDo", CreateTime: time.Now(), UpdateTime: time.Now(), DueDate: time.Now()},
			{ID: 2, Title: "Task 2", Status: "InProgress", CreateTime: time.Now(), UpdateTime: time.Now(), DueDate: time.Now()},
			{ID: 3, Title: "Task 3", Status: "Done", CreateTime: time.Now(), UpdateTime: time.Now(), DueDate: time.Now()},
		},
	}
}

func TestFindTasks(t *testing.T) {
	tasks := createSampleTaskList()

	// Mock user input
	mockInput := "ToDo\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(mockInput))
	r.Close()

	// Redirect stdout for capturing output
	oldStdout := os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	// Call the FindTasks function
	FindTasks(&tasks)

	// Read the output from stdout
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	// Restore stdin
	os.Stdin = os.NewFile(0, "/dev/stdin")

	// Check the output
	expectedOutput := "Task list:\n1. Task 1 (Status: ToDo, Due Date: " + tasks.Tasks[0].DueDate.Format("2006-01-02") +
		", Created At: " + tasks.Tasks[0].CreateTime.Format("2006-01-02 15:04:05") +
		", Updated At: " + tasks.Tasks[0].UpdateTime.Format("2006-01-02 15:04:05") + ")\n"
	if string(out) != expectedOutput {
		t.Errorf("Unexpected output:\nExpected: %s\nGot: %s", expectedOutput, string(out))
	}
}
