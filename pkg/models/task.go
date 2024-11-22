package models

// capital Task so it is exported, lowercase task would be private
type Task struct {
	ID        int
	Name      string
	Completed bool
}
