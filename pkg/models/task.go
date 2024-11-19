package models

// Task represents a task in the todo list
// capital Task so it is exported
// struct is a collection of fields
type Task struct {
	ID        int
	Name      string
	Completed bool
}
