package markdone

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/zlowram/gocli"
)

func (todo Todo) List(c gocli.Command) error {
	list := todo.tasks
	for i := 0; i < len(list); i++ {
		var title string
		if list[i].Completed {
			title = color.GreenString("%s", list[i].Title)
		} else {
			title = color.RedString("%s", list[i].Title)
		}
		fmt.Printf("%d. %s\n", (i + 1), title)
	}
	return nil
}
