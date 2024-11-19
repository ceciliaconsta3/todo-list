package todo

import (
	"fmt"
	"todo-list/pkg/models"
)

type TodoList struct {
	// Tasks is a slice of task instances
	Tasks []models.Task
}

// t is a pointer receiver reference to the task instance
// accepts taskname as a string
// t *TodoList is a method receiver w/in this method definition
// *TodoList is a pointer receiver, not a copy of TodoList instance so we can modify the original instance
// variable t is the reference, inside this method, use t to access or modify the fields in TodoList
// defines receiver type and specifies addTask is method withint TodoList - can all addTask on an instance of TodoList
// myTodolist.addtask()

func (t *TodoList) AddTask(taskName string) {
	// check if the task is not empty
	if taskName == "" {
		// if empty, return
		fmt.Println("Task cannot be empty")
		return
	}
	// assign a unique ID to the task
	// let's try length of the current list++
	newID := len(t.Tasks) + 1
	// create a new task instance
	newTask := models.Task{
		ID:        newID,
		Name:      taskName,
		Completed: false,
	}
	// Add a task to the tasks slice of the todolist
	// append is a built-in function to add an element to a slice
	// append(slice, element)
	t.Tasks = append(t.Tasks, newTask)
	// print a success message
	fmt.Println("Task: ", newTask.Name, "added successfully")
}

func (t *TodoList) ListTasks() {
	fmt.Println("Tasks: ", t.Tasks)
}

func (t *TodoList) DeleteTask(taskID int) {
	// check if task is empty
	if taskID == 0 {
		return
	}
	// check if taskID is within the range of the tasks slice
	// see if taskID exists
	// taskID-1 because this slice is 1 based not 0 based
	// the extra .ID access the ID field of the slice
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
	// check if task is empty
	if taskID <= 0 || taskID > len(t.Tasks) {
		fmt.Println("Task ID: ", taskID, " does not exist")
		return
	}
	// check if task isCompleted = false
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
