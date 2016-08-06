package markdone

import (
	"regexp"

	"github.com/zlowram/gocli"
)

func (todo Todo) ClearCompleted(c gocli.Command) error {
	expr := "(?m:(\\n)?^- \\[X\\] (.*)$)"
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
