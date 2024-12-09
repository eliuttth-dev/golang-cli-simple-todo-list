package main

import (
  "fmt"
  "os"
)

func saveTasksToFile(tasks []Task){
  file, err := os.Create("tasks.txt")

  if err != nil {
    fmt.Println("Error creating file:", err)
    return
  }

  defer file.Close()

  for _, task := range tasks {
    status := "Incomplete"
    if task.Completed{
      status = "Complete"
    }
    file.WriteString(fmt.Sprintf("%s [%s]\n", task.Description, status))
  }
}
