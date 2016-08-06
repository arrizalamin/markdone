package markdone

import (
	"strings"

	"github.com/zlowram/gocli"
)

func (todo Todo) Add(c gocli.Command) error {
	task := strings.Join(c.Flag.Args(), " ")
	text := todo.text + "\n- [ ] " + task
	return WriteFile(text)
}
