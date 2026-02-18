package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func (cfg *config) marshalTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal tasks file: %w", err)
	}

	err = os.WriteFile(cfg.TasksFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}
	return nil
}

func (cfg *config) getCurrentID() (int, error) {
	tasks, err := unmarshalTasks(cfg.TasksFilePath)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	maxID++
	return maxID, nil
}
