package help

import "fmt"

var helpMessage = "\nAction:\n\tview: view your notes\n\tadd: add a note\n\tdelete: delete a note\n\tclear: delete all note\n\thelp: show help message"

func Print() {
	fmt.Println(helpMessage)
}
