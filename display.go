package main

import "fmt"

func displayTasks(tasks []Task) {
  fmt.Println("To-Do List:")
  
  for i, task := range tasks {
    status := "Incomplete"
    if task.Completed {
      status = "Complete"
    }
    fmt.Printf("%d. %s [%s]\n", i + 1, task.Description, status)
  }
}
