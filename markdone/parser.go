package markdone

import (
	"io/ioutil"
	"regexp"
)

type Task struct {
	Title     string
	Completed bool
}

type Todo struct {
	text  string
	tasks []Task
}

func ParseTodos() (todo Todo, err error) {
	text, err := ioutil.ReadFile("TODOS.md")
	if err != nil {
		return
	}
	todo.text = string(text)
	r, err := regexp.Compile("(?m:^- \\[( |X)\\] (.*)$)")
	if err != nil {
		return
	}
	for _, line := range r.FindAllStringSubmatch(todo.text, -1) {
		task := Task{
			Title:     line[2],
			Completed: line[1] == "X",
		}
		todo.tasks = append(todo.tasks, task)
	}
	return
}
