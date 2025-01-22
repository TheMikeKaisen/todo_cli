# Todo CLI Application

A simple command-line application to manage your daily tasks. This application allows you to create, complete, delete, and list todos, while storing them in a JSON file. The project is written in Go and uses the `simpletable` package for a clean display of tasks.

---

## Features

- **Add**: Add new tasks to your todo list.
- **Complete**: Mark a task as completed.
- **Delete**: Remove a task from your list.
- **List**: View all your tasks in a formatted table.
- **Persistent Storage**: Saves todos in a JSON file (`.todo.json`) for persistence.

---

## Project Structure

```
├── .todo.json         # Hidden file to store todos (auto-generated)
├── Makefile           # Optional Makefile for build automation
├── cmd/
│   └── todo/
│       └── main.go    # Entry point of the application
├── colors.go          # Contains helper functions for colored output
├── go.mod             # Go module file
├── go.sum             # Go dependencies file
└── todo.go            # Core functionality of the Todo application

```

---

## Installation

1. Clone the repository:
    
    ```bash
    git clone https://github.com/TheMikeKaisen/todo_cli.git
    cd todo_cli
    
    ```
    
2. Build the application:
    
    ```bash
    go build -o todo ./cmd/todo/main.go
    
    ```
    
3. Verify the build:
    
    ```bash
    ./todo -h
    
    ```
    

---

## Usage

### Available Commands

- **Add a task:**
    
    ```bash
    ./todo -add Task description here
    
    ```
    
- **Complete a task:**
    
    ```bash
    ./todo -complete <task_number>
    
    ```
    
- **Delete a task:**
    
    ```bash
    ./todo -delete <task_number>
    
    ```
    
- **List all tasks:**
    
    ```bash
    ./todo -list
    
    ```
    

### Example

1. Add a task:
    
    ```bash
    ./todo -add Buy groceries
    ./todo -add Finish the Go project
    
    ```
    
2. List tasks:
    
    ```bash
    ./todo -list
    
    ```
    
3. Mark a task as completed:
    
    ```bash
    ./todo -complete=1
    
    ```
    
4. Delete a task:
    
    ```bash
    ./todo -delete=2
    
    ```
    

---

## Implementation Details

### `main.go`

- Handles command-line flags and parses user inputs.
- Calls appropriate methods from `todo.go` based on user commands.

### `todo.go`

- Core functionality for managing tasks.
- Implements methods for:
    - Adding tasks
    - Completing tasks
    - Deleting tasks
    - Saving and loading tasks from `.todo.json`
    - Printing tasks in a formatted table using `simpletable`.

### `colors.go`

- Helper functions to display colored output for better visual clarity.

---

## Dependencies

- **simpletable**: Used for displaying tasks in a clean table format.
    
    ```
    github.com/alexeyco/simpletable v1.0.0
    
    ```
    

---

## Permissions

If you encounter a `Permission Denied` error while running the built binary, ensure the file is executable:

```bash
chmod +x ./todo

```

---

## Error Handling

- If the `.todo.json` file is missing or corrupted, a new file will be created.
- Invalid indices for tasks (e.g., out of bounds) return appropriate error messages.
- Empty input for new tasks is rejected with an error.

---

## Future Enhancements

- Add support for task prioritization.
- Implement due dates and reminders.
- Improve error messages for better user feedback.
- Add unit tests to ensure code reliability.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

## Contribution

Contributions are welcome! Please feel free to submit a pull request or open an issue for suggestions and improvements.

---

## Author

[Karthik H](https://x.com/Karthik_h_nair)

Feel free to connect and share your feedback!