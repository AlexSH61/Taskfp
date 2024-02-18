package model

import (
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
