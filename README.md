# Task Management Program in Go

This is a simple command-line task management program written in Go. It allows users to add, remove, list, and mark tasks as complete through an interactive menu. The program stores tasks in an in-memory list and provides basic task operations.

## Features

- **List Tasks**: View all tasks with their respective indices.
- **Add Task**: Add a new task to the list.
- **Delete Task**: Remove a task by its index.
- **Mark Task as Complete**: Mark a task as finished by its index.
- **Exit**: Exit the application.

## Usage

To run this program, you'll need to have Go installed on your machine. After cloning or downloading the project, you can execute it using the following steps:

### Running the Program

1. **Install Go**: If you don't have Go installed, download and install it from the official Go website.
2. **Clone or Download the Code**: Save the source code into a file called `main.go`.
3. **Run the Program**: Open your terminal, navigate to the directory containing `main.go`, and run the following command:
   ```go
   go run main.go
   ```
4. **Follow the Interactive Menu**: The program will present you with a list of options. Simply input the corresponding number to perform the desired action.

## Example of Usage

When prompted, select an option by typing the number:

```
Please Input a Number Below
Option 1: List Tasks
Option 2: Add Task
Option 3: Delete Task
Option 4: Mark Task As Complete
Option 5: Exit
```

- If you choose option 2, you'll be asked to enter a task name:
  ```
  Input Name:
  ```
- If you choose option 1, the program will list all current tasks:
  ```
  0 Buy groceries
  1 Study Go
  ```
- If you choose option 4, you'll be asked to input the index of the task to mark as complete:
  ```
  What Index?
  ```

## Code Overview

The program consists of the following functions:

- `addTask()`: Prompts the user for a task name and adds it to the task list.
- `removeTask()`: Prompts the user for the index of the task to be removed and deletes it from the task list.
- `listTask()`: Displays all the tasks currently in the task list.
- `completeTask()`: Marks a task as complete based on its index.

## License

This project is free to use and distribute. Feel free to modify it to suit your needs.
