package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := []string{"add", "update", "delete", "mark-in-progress", "mark-done", "list", "list done", "list todo", "list in-progress"}

	cfg := &config{
		ID: 0,
		TasksFilePath: "tasks.json",
	}

	fmt.Println("Task-Tracker running ...")

	for scanner.Scan() {
		input := scanner.Text()

		// Split the input into fields and check if there are any commands
		inputFields := strings.Fields(input)
		if len(inputFields) < 1 {
			fmt.Println("Enter a command")
			continue
		}

		// Check if the first field is a valid command
		cmd, args := inputFields[0], inputFields[1:]
		if !slices.Contains(commands, cmd) {
			fmt.Printf("Invalid command\nCommands: %s\n", strings.Join(commands, ", "))
		}

		switch cmd {
		case "add":
			err := cfg.add(args)
			if err!= nil {
				fmt.Printf("Error adding task: %v\n", err)
			}
		default:
			fmt.Println("Command not implemented yet")
		}

	}

}
