package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Task struct {
	Finished bool `json:"Finished"`
	ToDo string `json:"ToDo"`
}

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

	newTask := Task{Finished: false, ToDo: taskName}

	data, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("Error Reading File:", err)
		return
	}

	var tasks_temp[] Task

	err2 := json.Unmarshal(data, &tasks_temp)

	if err2 != nil {
		fmt.Println("Error Marshalling tasks_temp:", err2)
		return
	}

	tasks_temp = append(tasks_temp, newTask)

	newJsonTask, err3 := json.Marshal(tasks_temp)

	if err3 != nil {
		fmt.Println("Error Marshalling tasks_temp:", err3)
		return
	}
	
	os.WriteFile("data.json", newJsonTask, 0644)

	/*newTask := Task{Finished: false, ToDo: taskName}

	tasks = append(tasks, newTask)

	newJsonTask, _ := json.Marshal(tasks)

	os.WriteFile("data.json", newJsonTask, 0644)*/
}

func removeTask() {
	var index int
	fmt.Println("Which index to remove?")
	fmt.Scan(&index)
	data, _ := os.ReadFile("data.json")

	var tasks_temp[] Task

	json.Unmarshal(data, &tasks_temp)

	if(index > len(tasks_temp) || index < 0) {
		fmt.Println("Make sure to input correct index")
		return
	}

	tasks_temp = append(tasks_temp[:index], tasks_temp[index + 1:]...)

	newJsonTask, err := json.Marshal(tasks_temp)

	if err != nil {
		fmt.Println("Error Marshalling tasks_temp:", err)
		return
	}

	os.WriteFile("data.json", newJsonTask, 0644)

}

func listTask() {

	data, _ := os.ReadFile("data.json")

	var tasks_temp[] Task

	err := json.Unmarshal(data, &tasks_temp)

	if err != nil {
		fmt.Println("Error Unmarshalling tasks_temp:", err)
		return
	}

	for i, task := range tasks_temp {
		fmt.Println(i, task.ToDo)
	}
}

func completeTask() {
	var index int
	fmt.Printf("What Index?: ")
	fmt.Scan(&index)
	data, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("Error Reading File:", err)
		return
	}

	var tasks_temp[] Task

	err2 := json.Unmarshal(data, &tasks_temp)

	if err2 != nil {
		fmt.Println("Error Unmarshalling tasks_temp:", err2)
		return
	}

	if(index > len(tasks_temp) || index < 0) {
		fmt.Println("Make sure to input correct index")
		return
	}

	tasks_temp[index].Finished = true

	newJsonTask, err3 := json.Marshal(tasks_temp)

	if err3 != nil {
		fmt.Println("Error Marshalling tasks_temp:", err3)
		return
	}

	os.WriteFile("data.json", newJsonTask, 0644)
}