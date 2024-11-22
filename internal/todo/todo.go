package todo

import (
	"fmt"
	"log"
	"os"
	"todo-list/pkg/models"
)

type TodoList struct {
	Tasks []models.Task
}

func (t *TodoList) AddTask(taskName string) {
	if taskName == "" {
		// if empty, return
		fmt.Println("Task cannot be empty")
		return
	}
	newID := len(t.Tasks) + 1
	newTask := models.Task{
		ID:        newID,
		Name:      taskName,
		Completed: false,
	}
	t.Tasks = append(t.Tasks, newTask)
	fmt.Println("Task: ", newTask.Name, "added successfully")
}

func (t *TodoList) ListTasks() {
	fmt.Println("Tasks: ", t.Tasks)
}

func (t *TodoList) DeleteTask(taskID int) {
	if taskID == 0 {
		return
	}
	// taskID-1 because this slice is 1 based not 0 based
	if taskID > 0 && taskID <= len(t.Tasks) && t.Tasks[taskID-1].ID == taskID {
		// search slice for task with ID
		// using range form of the for loop to iterate over a slice
		for index, task := range t.Tasks {
			if task.ID == taskID {
				// delete the task by appending the matching indexes into the append function directly
				// there is no built-in remove() func
				// so need to create a new slice with the elements before and after the element to be removed
				// t.Tasks[:index] = elements before
				// t.Task[index+1:] = elements after
				// index = the matching taskID
				t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
				// print success message
				fmt.Println("Task: ", t.Tasks[index].Name, "deleted successfully")
			}
		}
		// print out the entire list to see where we stand
		fmt.Println("Tasks: ", t.Tasks)
	}
}

func (t *TodoList) CompleteTask(taskID int) {
	if taskID <= 0 || taskID > len(t.Tasks) {
		fmt.Println("Task ID: ", taskID, " does not exist")
		return
	}
	// had to remove last condition because once something has been deleted the IDs no longer match their positions in the slice
	if taskID > 0 && taskID <= len(t.Tasks) {
		// search slice for task with ID
		// using range form of the for loop to iterate over a slice
		for index, task := range t.Tasks {
			if task.ID == taskID {
				// update task to completed = true
				t.Tasks[index].Completed = true
				// print success message
				fmt.Println("Task: ", t.Tasks[index].Name, "was completed successfully")
			}
		}
		// print out the entire list to see where we stand
		fmt.Println("Tasks: ", t.Tasks)
	}
}

func (t *TodoList) SaveTasksToFile() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks to save")
		return
	} else {
		file, err := os.OpenFile("todo-list.txt", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("There was a problem opening the file: ", err)
		}
		defer file.Close()
		for _, task := range t.Tasks {
			taskStr := fmt.Sprintf("ID: %d, Name: %s, Completed: %v", task.ID, task.Name, task.Completed)
			_, err := file.WriteString(taskStr + "\n")
			if err != nil {
				log.Fatal("There was a problem saving the file: ", err)
			}
		}
		fmt.Println("All tasks saved successfully")
		fmt.Println("Todo List saved to todo-list.txt")
	}
}
