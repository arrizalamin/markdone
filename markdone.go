package main

import (
	"flag"
	"log"

	"github.com/arrizalamin/markdone/markdone"
	"github.com/zlowram/gocli"
)

func main() {
	flag.Parse()

	todo, err := markdone.ParseTodos()
	if err != nil {
		log.Fatalln(err)
	}

	c := gocli.NewCli("todo", "Manage Todo List", flag.Args())
	c.Commands = []gocli.Command{
		// List
		gocli.Command{
			Name:        "list",
			ShortName:   "ls",
			Run:         todo.List,
			Description: "Print all pending todo items",
		},

		// Add Task
		gocli.Command{
			Name:        "add",
			Run:         todo.Add,
			Description: "Create a new task",
		},

		// Remove Task
		gocli.Command{
			Name:        "remove",
			ShortName:   "rm",
			Run:         todo.Remove,
			Description: "Remove a task",
		},

		// Complete Task
		gocli.Command{
			Name:        "complete",
			ShortName:   "do",
			Run:         todo.Complete,
			Description: "Mark a task as completed",
		},

		// Uncomplete Task
		gocli.Command{
			Name:        "revert",
			ShortName:   "undo",
			Run:         todo.Uncomplete,
			Description: "Revert a task to incomplete",
		},

		// Clear Completed Task
		gocli.Command{
			Name:        "clear-completed",
			Run:         todo.ClearCompleted,
			Description: "Remove all completed tasks",
		},
	}
	if err := c.Handle(); err != nil {
		log.Println(err)
	}
}
