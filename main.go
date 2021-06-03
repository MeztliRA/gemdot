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
	"github.com/MeztliRA/gemdot/config"
	c "github.com/MeztliRA/gemdot/constants"
	"github.com/MeztliRA/gemdot/file"
	u "github.com/MeztliRA/gemdot/utils"
	cfg "github.com/olebedev/config"
)

func init() {
	log.SetPrefix("gemdot: ")
	log.SetFlags(0)
}

func main() {
	file.Check()

	notes := file.Read()
	config := config.Read()

	userAction(notes, config)
}

func userAction(notes []string, config *cfg.Config) {
	reader := bufio.NewReader(os.Stdin)

	if greet, err := config.Bool("greeting"); greet && err == nil {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		username := user.Username
		color.Greenf("hello, %s!\n", username)
	} else if err != nil {
		log.Fatal(err)
	}

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
		response = strings.TrimSpace(response)

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
