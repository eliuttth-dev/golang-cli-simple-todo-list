package main

import "fmt"

func completeTask(tasks []Task, taskIndex int) []Task {
  if taskIndex >= 0 && taskIndex < len(tasks){
    tasks[taskIndex].Completed = true
  } else {
    fmt.Println("Invalid task index")
  }
  return tasks
}
