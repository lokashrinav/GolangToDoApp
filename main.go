package main

import (
	"fmt"
)

type Task struct {
	finished bool
	toDo string
}

var tasks[] Task

func main() {
	var name string;
	for {
		fmt.Println("Please Input a Number Below")
		fmt.Println("Option 1: List Tasks")
		fmt.Println("Option 2: Add Task")
		fmt.Println("Option 3: Delete Task")
		fmt.Println("Option 4: Mark Task As Complete")
		fmt.Println("Option 5: Exit")
		fmt.Scan(&name)
		if name == "1" {
			listTask()
		} else if name == "2" {
			addTask()
		} else if name == "3" {
			removeTask()
		} else if name == "4" {
			completeTask()
		} else if name == "5" {
			break
		}
		fmt.Println("")
	}
}

func addTask() {
	var taskName string
	fmt.Println("Input Name: ")
	fmt.Scan(&taskName)
	tasks = append(tasks, Task{finished: false, toDo: taskName})
}

func removeTask() {
	var index int
	fmt.Println("Which index to remove?")
	fmt.Scan(&index)
	tasks = append(tasks[:index], tasks[index + 1:]...)
}

func listTask() {
	for i, task := range tasks {
		fmt.Println(i, task.toDo)
	}
}

func completeTask() {
	var index int
	fmt.Printf("What Index?: ")
	fmt.Scan(&index)
	tasks[index].finished = true
}