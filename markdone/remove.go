package markdone

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/zlowram/gocli"
)

func (todo Todo) Remove(c gocli.Command) error {
	index, err := strconv.Atoi(c.Flag.Arg(0))
	if err != nil {
		return err
	}
	list := todo.tasks
	if index > len(list) {
		return errors.New("Task not found")
	}
	expr := "(?m:(\\n)?^- \\[(?: |X)\\] " + list[index-1].Title + "$)"
	r, err := regexp.Compile(expr)
	if err != nil {
		return err
	}
	text := r.ReplaceAllString(todo.text, "")
	if err = WriteFile(text); err != nil {
		return err
	}
	return nil
}
