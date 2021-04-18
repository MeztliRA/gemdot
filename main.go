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

func init() {
	log.SetPrefix("gemdot: ")
	log.SetFlags(0)
}

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

	firstTime := true
	for {
		if firstTime {
			color.Green("what do you want to do(view, add, delete, clear, help, quit) ")
		} else {
			color.Green("\nwhat do you want to do(view, add, delete, clear, help, quit) ")
		}
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.Trim(response, "\n")

		switch response {
		case "View", "view", "VIEW":
			files.View(notes)
		case "Add", "add", "ADD":
			notes = files.Add(notes)
			files.Overwrite(notes)
		case "Delete", "delete", "DELETE":
			notes = files.Delete(notes)
			files.Overwrite(notes)
		case "Clear", "clear", "CLEAR":
			notesGet, cleared := files.Clear()
			if cleared {
				files.Overwrite(notesGet)
			}
		case "Help", "help", "HELP":
			help.Print()
		case "Quit", "quit", "QUIT":
			os.Exit(0)
		default:
			color.Red("unknown action")
		}

		firstTime = false
	}
}
