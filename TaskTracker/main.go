package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	ColorReset = "\033[0m"
	ColorRed = "\033[31m"
	ColorGreen = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue = "\033[34m"
	ColorGray = "\033[37m"
)

type Task struct {
	ID int `json: "id"`
	Description string `json: "description"`
	Status string `json: "status"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "UpdatedAt"`
}

type Tasks []Task

const (
	fileName = "tasks.json"
	statusTodo = "todo"
	statusInProgress = "in-progress"
	statusDone = "done"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	markInProgressCmd := flag.NewFlagSet("mark-in-progress", flag.ExitOnError)
	markDoneCmd := flag.NewFlagSet("mark-done", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	addDescription := addCmd.String("description", "", "Task description")

	updateID := updateCmd.Int("id", 0, "Task ID")
	updateDescription := updateCmd.String("description", "", "New task description")

	deleteID := deleteCmd.Int("id", 0, "Task ID")

	markInProgressID := markInProgressCmd.Int("id", 0, "Task ID")
	markDoneID := markDoneCmd.Int("id", 0, "Task ID")

	listStatus := listCmd.String("status", "", "Filter by status: todo, in-progress, done")

	if len(os.Args) < 2 {
		printInstruction()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addDescription == "" {
			fmt.Printf("%sError: Description is required$s\n", ColorRed, ColorReset)
			addCmd.Usage()
			os.Exit(1)
		}
		addTask(*addDescription)
	
	case "update":
		updateCmd.Parse(os.Args[2:0])
		if *updateID <= 0 || *updateDescription == "" {
			
		}
	}
}

func printInstruction() {
	fmt.Println(` Usage: task-cli <command> [arguments]
	Commands:
		add "description" 					Add a new task
		update <id> "new description"		Update a task
		delete <id> 						Delete a task
		mark-in-progress <id> 				Mark task as in progress
		mark-done <id>						Mark task as done
		list [status]						List tasks(optional: todo, in-progress, done)

	Examples:
		task-cli add "Buy groceries"
		task-cli update 1 "Buy groceries and cook"
		task-cli delete 1
		task-cll mark-in-progress 1
		task-cli mark-done 1
	`)
}

