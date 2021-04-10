package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	files "github.com/MeztliRA/gemdot/file"
	note "github.com/MeztliRA/gemdot/notes"
)

func main() {
	checkFile()

	notes := readFile()

	reader := bufio.NewReader(os.Stdin)

L:
	for {
		fmt.Print("what do you want to do(view, add, delete, clear) ")
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
		default:
			fmt.Println("unknown action")
			continue
		}
	}
}

func checkFile() {
	if _, err := os.Stat(note.File); os.IsNotExist(err) {
		if _, err := os.Stat(note.Directory); os.IsNotExist(err) {
			dirErr := os.Mkdir(note.Directory, 0755)
			if dirErr != nil {
				log.Fatal(dirErr)
			}
		}

		var notes []string

		files.Overwrite(notes)
	}
}

func readFile() []string {
	file, err := os.ReadFile(note.File)
	if err != nil {
		log.Fatal(err)
	}

	var notes []string

	unmarshalErr := json.Unmarshal(file, &notes)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}

	return notes
}
