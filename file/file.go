package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	note "github.com/MeztliRA/gemdot/notes"
	"github.com/MeztliRA/yon"
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

	fmt.Print("\nenter new note: ")
	note, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	note = strings.Trim(note, "\n")

	notes = append(notes, note)

	fmt.Println("\nadded new note!")

	return notes
}

func Delete(notes []string) []string {
	fmt.Println("\nnotes:")
	if len(notes) == 0 {
		fmt.Println("\tno notes to delete")
		return notes
	} else {
		reader := bufio.NewReader(os.Stdin)

		for i, v := range notes {
			fmt.Printf("\t[%d] %s\n", i, v)
		}

		for {
			fmt.Print("\nplease enter the id of the note you want to delete: ")
			inputtedId, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			inputtedId = strings.Trim(inputtedId, "\n")
			id, err := strconv.Atoi(inputtedId)
			if err != nil {
				continue
			}

			if id >= len(notes) {
				continue
			}

			if id < 0 {
				continue
			}

			fmt.Println("\ndeleting note...")
			notes = removeIndex(notes, id)
			fmt.Println("\ndone!")

			return notes
		}
	}
}

func Clear() ([]string, bool) {
	var (
		notes   []string
		cleared bool
	)

	response := yon.Prompt("\nare you sure you want to delete all your note")
	if response == yon.Yes {
		fmt.Println("\nall notes deleted!")
		cleared = true
		return notes, cleared
	} else {
		fmt.Println("\ncancelled...")
		cleared = false
		return notes, cleared
	}
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

func removeIndex(notes []string, index int) []string {
	return append(notes[:index], notes[index+1:]...)
}
