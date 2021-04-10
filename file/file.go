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
	"github.com/fatih/color"
)

var (
	magenta = color.New(color.FgHiMagenta).PrintlnFunc()
	green   = color.New(color.FgGreen).PrintFunc()
	hiGreen = color.New(color.FgHiGreen).PrintlnFunc()
	red     = color.New(color.FgHiRed).PrintlnFunc()
)

func View(notes []string) {
	magenta("\nnotes:")
	if len(notes) == 0 {
		fmt.Println("\tno notes")
	} else {
		for _, v := range notes {
			fmt.Printf("\t- %s\n", v)
		}
	}
}

func Add(notes []string) []string {
	reader := bufio.NewReader(os.Stdin)

	green("\nenter new note: ")
	note, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	note = strings.Trim(note, "\n")

	notes = append(notes, note)

	hiGreen("\nadded new note!")

	return notes
}

func Delete(notes []string) []string {
	magenta("\nnotes:")
	if len(notes) == 0 {
		fmt.Println("\tno notes to delete")
		return notes
	}

	reader := bufio.NewReader(os.Stdin)
	blue := color.New(color.FgBlue).PrintfFunc()

	for i, v := range notes {
		blue("\t[%d] ", i)
		fmt.Printf("%s\n", v)
	}

	for {
		green("\nplease enter the id of the note you want to delete: ")
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

		hiGreen("\ndeleting note...")
		notes = removeIndex(notes, id)
		hiGreen("\ndone!")

		return notes
	}
}

func Clear() ([]string, bool) {
	var (
		notes   []string
		cleared bool
	)

	color.Set(color.FgGreen)
	response := yon.Prompt("\nare you sure you want to delete all your note")
	color.Unset()
	if response == yon.Yes {
		hiGreen("\nall notes deleted!")
		cleared = true
		return notes, cleared
	} else {
		red("\ncancelled...")
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
