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

	"github.com/MeztliRA/gemdot/color"
	c "github.com/MeztliRA/gemdot/constants"
	"github.com/MeztliRA/gemdot/file"
	u "github.com/MeztliRA/gemdot/utils"
)

func init() {
	log.SetPrefix("gemdot: ")
	log.SetFlags(0)
}

func main() {
	file.Check()

	notes := file.Read()

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
			color.Green(c.ActionMessage)
		} else {
			color.Green("\n" + c.ActionMessage)
		}
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = u.TrimString(response)

		switch response {
		case "View", "view", "VIEW":
			file.View(notes)
		case "Add", "add", "ADD":
			notes = file.Add(notes)
			file.Overwrite(notes)
		case "Delete", "delete", "DELETE":
			notes = file.Delete(notes)
			file.Overwrite(notes)
		case "Clear", "clear", "CLEAR":
			notesGet, cleared := file.Clear()
			if cleared {
				file.Overwrite(notesGet)
			}
		case "Version", "version", "VERSION":
			u.PrintVersion()
		case "Help", "help", "HELP":
			u.PrintHelp()
		case "Quit", "quit", "QUIT":
			os.Exit(0)
		default:
			color.Red("unknown action")
		}

		firstTime = false
	}
}
