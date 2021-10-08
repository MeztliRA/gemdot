/*
A CLI app to store your notes.

written in go.
*/
package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/MeztliRA/gemdot/color"
	"github.com/MeztliRA/gemdot/config"
	c "github.com/MeztliRA/gemdot/constants"
	"github.com/MeztliRA/gemdot/note"
	u "github.com/MeztliRA/gemdot/utils"
	fc "github.com/fatih/color"
	cfg "github.com/olebedev/config"
)

func init() {
	log.SetPrefix("gemdot: ")
	log.SetFlags(0)
}

func main() {
	note.Check()
	config.Check()

	notes := note.Read()
	config := config.Read()

	userCommand(notes, config)
}

func userCommand(notes []string, config *cfg.Config) {
	reader := bufio.NewReader(os.Stdin)

	if noColor, err := config.Bool("no-color"); noColor && err == nil {
		fc.NoColor = true
	} else if err != nil {
		fc.NoColor = false
	} else {
		fc.NoColor = false
	}

	if greet, err := config.Bool("greeting"); greet && err == nil || err != nil {
		u.PrintGreeting()
	}

	firstTime := true
	for {
		if firstTime {
			color.Green(c.CommandMessage)
		} else {
			color.Green("\n" + c.CommandMessage)
		}
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.TrimSpace(response)

		switch response {
		case "View", "view", "VIEW":
			note.View(notes)
		case "Add", "add", "ADD":
			notes = note.Add(notes)
			note.Overwrite(notes)
		case "Delete", "delete", "DELETE":
			notes = note.Delete(notes)
			note.Overwrite(notes)
		case "Clear", "clear", "CLEAR":
			notesGet, cleared := note.Clear()
			if cleared {
				note.Overwrite(notesGet)
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
