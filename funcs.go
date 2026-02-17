package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func (cfg *config) add(args []string) error {
	// Check if the description argument is provided
	if len(args) < 1 {
		fmt.Println("Usage: add <description>")
		return nil
	}

	description := args[0]

	tasks, err := unmarshalTasks(cfg.TasksFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	t := Task{
		ID:          cfg.ID,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, t)

	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal tasks file: %w", err)
	}

	err = os.WriteFile(cfg.TasksFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}

	fmt.Printf("Task added: %s (ID: %v)\n", description, t.ID)

	// Increment the ID for the next task
	cfg.ID++

	return nil
}

func unmarshalTasks(path string) ([]Task, error) {
	tasks := []Task{}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return tasks, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to check tasks file: %w", err)
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}

	if len(f) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(f, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return tasks, nil
}
