/*
A CLI app to store your notes.

written in go.
*/
package main

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/MeztliRA/gemdot/color"
	files "github.com/MeztliRA/gemdot/file"
	"github.com/MeztliRA/gemdot/help"
)

func main() {
	files.Check()

	notes := files.Read()

	userAction(notes)
}

func userAction(notes []string) {
	reader := bufio.NewReader(os.Stdin)

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	username := user.Username
	color.Greenf("hello, %s!\n", username)
L:
	for {
		color.Green("what do you want to do(view, add, delete, clear, help) ")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.Trim(response, "\n")

		switch response {
		case "View", "view", "VIEW":
			files.View(notes)
			break L
		case "Add", "add", "ADD":
			notes = files.Add(notes)
			files.Overwrite(notes)
			break L
		case "Delete", "delete", "DELETE":
			notes = files.Delete(notes)
			files.Overwrite(notes)
			break L
		case "Clear", "clear", "CLEAR":
			notesGet, cleared := files.Clear()
			if cleared {
				files.Overwrite(notesGet)
			}
			break L
		case "Help", "help", "HELP":
			help.Print()
			break L
		default:
			color.Red("unknown action")
			continue
		}
	}
}
