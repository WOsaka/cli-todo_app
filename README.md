# CLI Todo App

A simple interactive command-line task tracker written in Go. Tasks are persisted to a local `tasks.json` file with no external dependencies.

## Requirements

- Go 1.25.1 or later

## Installation

```bash
git clone https://github.com/WOsaka/cli-todo_app
cd cli-todo_app
go build -o task-tracker
```

## Usage

Run the application:

```bash
./task-tracker
```

The app starts an interactive REPL loop. Type commands and press Enter to execute them.

```
Task-Tracker running ...
add Buy groceries
Task added: Buy groceries (ID: 1)
```

## Commands

| Command | Usage | Description |
|---|---|---|
| `add` | `add <description>` | Add a new task with status `todo` |
| `update` | `update <ID> <description>` | Update the description of an existing task |
| `delete` | `delete <ID>` | Delete a task by ID |
| `mark-in-progress` | `mark-in-progress <ID>` | Set a task's status to `in-progress` |
| `mark-done` | `mark-done <ID>` | Set a task's status to `done` |
| `list` | `list` | List all tasks |
| `list` | `list todo` | List tasks with status `todo` |
| `list` | `list in-progress` | List tasks with status `in-progress` |
| `list` | `list done` | List tasks with status `done` |

## Task Structure

Each task stored in `tasks.json` has the following fields:

| Field | Type | Description |
|---|---|---|
| `ID` | int | Auto-incremented unique identifier |
| `Description` | string | Task description text |
| `Status` | string | One of: `todo`, `in-progress`, `done` |
| `CreatedAt` | time.Time | Timestamp when the task was created |
| `UpdatedAt` | time.Time | Timestamp of the last update |

## Example Session

```
Task-Tracker running ...
add Write unit tests
Task added: Write unit tests (ID: 1)
add Fix login bug
Task added: Fix login bug (ID: 2)
list
ID      Description
1       Write unit tests
2       Fix login bug
mark-in-progress 1
Status of task 1 changed to: "in-progress"
mark-done 2
Status of task 2 changed to: "done"
list in-progress
ID      Description
1       Write unit tests
update 1 Write and run unit tests
Description of task 1 updated to: Write and run unit tests
delete 2
Task 2 deleted
```

## Project Structure

```
.
├── main.go       # Entry point and REPL loop
├── commands.go   # Command implementations (add, update, delete, list, mark-*)
├── funcs.go      # Helper functions for JSON serialization and ID management
├── structs.go    # Data types: Task and config
├── tasks.json    # Persisted task data (auto-created on first use)
└── go.mod        # Go module definition
```
