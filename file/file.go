package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	note "github.com/MeztliRA/gemdot/notes"
)

func View(notes []string) {
	fmt.Println("\nnotes:")
	if len(notes) == 0 {
		fmt.Println("\tno notes")
	} else {
		for i := 0; i < len(notes); i++ {
			fmt.Printf("\t- %s\n", notes[i])
		}
	}
}

func Add(notes []string) []string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter new note: ")
	note, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	note = strings.Trim(note, "\n")

	notes = append(notes, note)

	fmt.Println("added new note!")

	return notes
}

func Overwrite(notes []string) {
	jsonData, err := json.MarshalIndent(notes, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	writeErr := os.WriteFile(note.File, jsonData, 0644)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
