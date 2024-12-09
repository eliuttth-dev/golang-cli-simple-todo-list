package main

import "fmt"

func removeTask(tasks []Task, taskIndex int) []Task{
  if taskIndex >= 0 && taskIndex < len(tasks){
    tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
  } else {
    fmt.Println("Invalid Task index")
  }
  return tasks
}

