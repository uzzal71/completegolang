package main

import "fmt"

type TodoList struct {
	tasks []Task
}

func (td *TodoList) AddTask(taskName string) {
	fmt.Println("Task Added: ", taskName)
}

func (td *TodoList) ViewTask() {
	fmt.Println("Tasks are", td.tasks)
}

func (td *TodoList) MarkTaskAsDone(taskId int) {
    fmt.Println("Task marked as done", taskId)
}