package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  var tasks []Task

  reader := bufio.NewReader(os.Stdin)
  
  for{
    fmt.Println("\nTo-Do List CLI")
    fmt.Println("1. Add Task")
    fmt.Println("2. Remove Task")
    fmt.Println("3. Display Tasks")
    fmt.Println("4. Mark Task as Complete")
    fmt.Println("5. Save and Exit")
    fmt.Println("Choose an option: ")

    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    switch choice {
      case "1":
        fmt.Print("Enter task description: ")
        description, _ := reader.ReadString('\n')
        description = strings.TrimSpace(description)
        tasks = addTask(tasks, description)
      case "2":
        displayTasks(tasks)
        fmt.Print("Enter task number to remove: ")
        var taskIndex int
        fmt.Scanf("%d", &taskIndex)
        tasks = removeTask(tasks, taskIndex-1)
      case "3":
        displayTasks(tasks)
      case "4":
        displayTasks(tasks)
        fmt.Print("Enter task number to mark as complete: ")
        var taskIndex int
        fmt.Scanf("%d", &taskIndex)
        tasks = completeTask(tasks, taskIndex-1)
      case "5":
        saveTasksToFile(tasks)
        fmt.Println("Tasks saved to tasks.txt Exiting...")
        return
      default:
        fmt.Println("Invalid choice. Please try again")
    }
  }
}
