package help

import (
	"fmt"

	"github.com/MeztliRA/gemdot/color"
)

var (
	helpMessageHeader = "\nAction:"
	helpMessage       = "\n\tview: view your notes\n\tadd: add a note\n\tdelete: delete a note\n\tclear: delete all note\n\thelp: show help message\n\tquit: quit the app"
)

func Print() {
	color.Magenta(helpMessageHeader)
	fmt.Println(helpMessage)
}
