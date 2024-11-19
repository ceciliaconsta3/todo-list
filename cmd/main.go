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
	// create user prompts / instructions
	fmt.Println("Welcome to the Todo List App")
	fmt.Println("Enter a task to add to your list")
	fmt.Println("Enter 'delete' to remove a task")
	fmt.Println("Enter 'complete' to mark a task as complete")
	fmt.Println("Enter 'list' to see all tasks")
	fmt.Println("Enter 'exit' to quit the app")

	// forgot to instanciate the TodoList struct
	todoList := todo.TodoList{}

	// create a loop to keep the app running as long as the user doesn't input exit
	// that's silly, just have exit, exit the app not put as a condition in the loop
	for {
		// moved inside loop to make sure program doesn't end after user input
		// create Scanner instance to receive user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// handler user input
		// since expecting multople values let's use a switch statement
		// use default as catchall for the actual task strings
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
			// how to convert this string to an int?
			// use the strconv package
			// what is atoi?
			// it converts a string to an integer
			// returns 2 things: string and error - need to account for error
			iD, err := strconv.Atoi(taskID)
			if err != nil {
				fmt.Println("No task to complete")
				continue
			}
			todoList.CompleteTask(iD)
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

	// collect user input
	// call the AddTask method on the todoList instance

	// TEST DATA
	// call the AddTask method on the todoList instance
	// todoList.AddTask("Wash hair")
	// todoList.AddTask("Detangle hair")
	// todoList.AddTask("Apply leave-in conditioner")
	// todoList.AddTask("Seal with oil")
	// todoList.DeleteTask(1)
	// todoList.CompleteTask(3)
}
