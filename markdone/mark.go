package markdone

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/zlowram/gocli"
)

func mark(complete bool, c gocli.Command, todo Todo) error {
	index, err := strconv.Atoi(c.Flag.Arg(0))
	if err != nil {
		return err
	}
	list := todo.tasks
	if index > len(list) {
		return errors.New("Task not found")
	}
	title := list[index-1].Title
	expr := "(?m:(\\n)?^- \\[(?: |X)\\] " + title + ")"
	r, err := regexp.Compile(expr)
	if err != nil {
		return err
	}
	mark := " "
	if complete {
		mark = "X"
	}
	text := r.ReplaceAllString(todo.text, "\n- ["+mark+"] "+title)
	if err = WriteFile(text); err != nil {
		return err
	}
	return nil
}

func (todo Todo) Complete(c gocli.Command) error {
	return mark(true, c, todo)
}

func (todo Todo) Uncomplete(c gocli.Command) error {
	return mark(false, c, todo)
}
