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
	if _, err := os.Stat(note.File); os.IsNotExist(err) {
		if _, err := os.Stat(note.Directory); os.IsNotExist(err) {
			dirErr := os.Mkdir(note.Directory, 0755)
			if dirErr != nil {
				log.Fatal(dirErr)
			}
		}

		var notes []string

		jsonData, err := json.MarshalIndent(notes, "", "	")
		if err != nil {
			log.Fatal(err)
		}

		writeErr := os.WriteFile(note.File, jsonData, 0644)
		if writeErr != nil {
			log.Fatal(writeErr)
		}
	}

	file, err := os.ReadFile(note.File)
	if err != nil {
		log.Fatal(err)
	}

	var notes []string

	unmarshalErr := json.Unmarshal(file, &notes)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}

	reader := bufio.NewReader(os.Stdin)

L:
	for {
		fmt.Print("what do you want to do(view, add) ")
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
		default:
			fmt.Println("unknown action")
			continue
		}
	}
}