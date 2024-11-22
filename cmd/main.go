package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/internal/todo"
)

func main() {
	fmt.Println("Welcome to the Todo List App")
	fmt.Println("Enter a task to add to your list")
	fmt.Println("Enter 'delete' to remove a task")
	fmt.Println("Enter 'complete' to mark a task as complete")
	fmt.Println("Enter 'list' to see all tasks")
	fmt.Println("Enter 'exit' to quit the app")
	fmt.Println("Enter 'save' to save your tasks to a file")

	todoList := todo.TodoList{}

	for {
		// moved inside loop to make sure program doesn't end after user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// since expecting multiple values let's use a switch statement
		switch input {
		case "delete":
			fmt.Println("Enter the task ID to delete")
			scanner.Scan()
			taskID := strings.TrimSpace(scanner.Text())
			// how to convert this string to an int?
			// use the strconv package
			// what is atoi?
			// it converts a string to an integer
			// returns 2 things: string and error - need to account for error
			iD, err := strconv.Atoi(taskID)
			if err != nil {
				fmt.Println("Can't find that task ID")
				continue
			}
			todoList.DeleteTask(iD)
		case "complete":
			fmt.Println("Enter the task ID to complete")
			scanner.Scan()
			taskID := strings.TrimSpace(scanner.Text())
			iD, err := strconv.Atoi(taskID)
			if err != nil {
				fmt.Println("No task to complete")
				continue
			}
			todoList.CompleteTask(iD)
		case "save":
			fmt.Println("Tasks successfully saved to file")
			todoList.SaveTasksToFile()
			continue
		case "list":
			fmt.Println("Here are all the tasks")
			todoList.ListTasks()
			continue
		case "exit":
			fmt.Println("Exiting the app")
			os.Exit(0)
		default:
			todoList.AddTask(input)
		}
	}
}
