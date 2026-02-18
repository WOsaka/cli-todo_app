package main

import (
	"fmt"
	"strconv"
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
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, t)

	err = cfg.marshalTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task added: %s (ID: %v)\n", description, t.ID)

	// Increment the ID for the next task
	cfg.ID++

	return nil
}

func (cfg *config) update(args []string) error {
	// Check if the description argument is provided
	if len(args) < 2 {
		fmt.Println("Usage: update <ID> <description>")
		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	description := args[1]

	tasks, err := unmarshalTasks(cfg.TasksFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			break
		}
		if i == len(tasks)-1 {
			fmt.Println("Task not found")
			return nil
		}
	}

	err = cfg.marshalTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Description of task %v updated to: %s\n", id, description)
	return nil
}

func (cfg *config) delete(args []string) error {
	// Check if the description argument is provided
	if len(args) < 1 {
		fmt.Println("Usage: delete <ID>")
		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}

	tasks, err := unmarshalTasks(cfg.TasksFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("Task not found")
		return nil
	}

	for i, task := range tasks {
		if task.ID == id {
			fmt.Println(len(tasks))
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}

		if i == len(tasks)-1 {
			fmt.Println("Task not found")
			return nil
		}
	}

	err = cfg.marshalTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task %v deleted\n", id)
	return nil
}

func (cfg *config) markInProgress(args []string) error {
	// Check if the description argument is provided
	if len(args) < 1 {
		fmt.Println("Usage: mark-in-progress <ID>")
		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}

	tasks, err := unmarshalTasks(cfg.TasksFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("Task not found")
		return nil
	}

	for i, task := range tasks {
		if task.ID == id {
			if task.Status == "in-progress" {
				fmt.Printf("Task %v is already \"in-progress\"\n", id)
				return nil
			}

			tasks[i].Status = "in-progress"
			break
		}

		if i == len(tasks)-1 {
			fmt.Println("Task not found")
			return nil
		}
	}

	err = cfg.marshalTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Status of task %v changed to: \"in-progress\"\n", id)
	return nil
}

func (cfg *config) markDone(args []string) error {
	// Check if the description argument is provided
	if len(args) < 1 {
		fmt.Println("Usage: mark-done <ID>")
		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}

	tasks, err := unmarshalTasks(cfg.TasksFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("Task not found")
		return nil
	}

	for i, task := range tasks {
		if task.ID == id {
			if task.Status == "done" {
				fmt.Printf("Task %v is already \"done\"\n", id)
				return nil
			}

			tasks[i].Status = "done"
			break
		}

		if i == len(tasks)-1 {
			fmt.Println("Task not found")
			return nil
		}
	}

	err = cfg.marshalTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Status of task %v changed to: \"done\"\n", id)
	return nil
}
