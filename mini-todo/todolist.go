package main

import "fmt"

type TodoList struct {
	tasks []Task
}

func (td *TodoList) AddTask(taskName string) {
	taskId := len(td.tasks) + 1
	task := NewTask(taskId, taskName)
	td.tasks = append(td.tasks, task)
	fmt.Println("Task Added Successfully!")
}

func (td *TodoList) ViewTask() {
	fmt.Println("======================== TASKS ========================")
	for i, task := range td.tasks {
		doneStatus := " "
		if task.Done {
            doneStatus = "X"
        }
        fmt.Printf("[%s] Task #%d: %s\n", doneStatus, i + 1, task.Name)
    }
	fmt.Println("========================================================")
}

func (td *TodoList) MarkTaskAsDone(taskId int) {
    if taskId < 1 || taskId > len(td.tasks) {
		fmt.Println("Invalid Task ID")
	}
	td.tasks[taskId-1].MaskAsDone()
}