package main

func addTask(tasks []Task, description string ) []Task {
  task := Task{Description: description, Completed: false}
  return append(tasks, task)
}
