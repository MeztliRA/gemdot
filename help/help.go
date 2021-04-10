package help

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	magenta           = color.New(color.FgHiMagenta).PrintFunc()
	helpMessageHeader = "\nAction:"
	helpMessage       = "\n\tview: view your notes\n\tadd: add a note\n\tdelete: delete a note\n\tclear: delete all note\n\thelp: show help message"
)

func Print() {
	magenta(helpMessageHeader)
	fmt.Println(helpMessage)
}
