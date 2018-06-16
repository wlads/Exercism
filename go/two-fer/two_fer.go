package twofer

import "fmt"

const textTemplate = `One for %s, one for me.`

func ShareWith(name string) string {
	if len(name) == 0 {
		return fmt.Sprintf(textTemplate, "you")
	} else {
		return fmt.Sprintf(textTemplate, name)
	}
	
}
