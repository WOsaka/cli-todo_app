package main

import "time"

type config struct {
	ID            int
	TasksFilePath string
}

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
