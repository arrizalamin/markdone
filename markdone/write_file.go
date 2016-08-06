package markdone

import (
	"io/ioutil"
)

func WriteFile(text string) error {
	return ioutil.WriteFile("TODOS.md", []byte(text), 0644)
}
