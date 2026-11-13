package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Task is an in-memory todo task.
type Task struct {
	ID    int
	Title string
	Done  bool
}

// Store is a simple in-memory task store.
type Store struct {
	tasks []Task
	next  int
}

func NewStore() *Store {
	return &Store{next: 1}
}

func (s *Store) Add(title string) (Task, error) {
	// TODO: reject empty titles
	return Task{}, nil
}

func (s *Store) List() []Task {
	// TODO: return a copy
	return nil
}

func (s *Store) Complete(id int) error {
	// TODO
	return nil
}

func (s *Store) Delete(id int) error {
	// TODO
	return nil
}

func main() {
	store := NewStore()
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Todo CLI — commands: add <title>, list, done <id>, delete <id>, exit")

	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}

		line := strings.TrimSpace(reader.Text())
		parts := strings.SplitN(line, " ", 2)
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			// TODO: call store.Add
		case "list":
			// TODO: print tasks
		case "done":
			// TODO: parse ID and complete task
		case "delete":
			// TODO: parse ID and delete task
		case "exit":
			return
		default:
			fmt.Println("unknown command")
		}

		if len(parts) == 0 {
			fmt.Println(errors.New("empty command"))
		}
	}
}
