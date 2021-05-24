package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/MeztliRA/gemdot/color"
	"github.com/MeztliRA/gemdot/note"
	u "github.com/MeztliRA/gemdot/utils"
	"github.com/MeztliRA/yon"
)

func View(notes []string) {
	color.Magentaln("\nnotes:")
	if len(notes) == 0 {
		color.Red("\tno notes")
	} else {
		for _, v := range notes {
			fmt.Printf("\t• %s\n", v)
		}
	}
}

func Add(notes []string) []string {
	reader := bufio.NewReader(os.Stdin)

	color.Magenta("\nenter new note: ")
	note, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	note = strings.Trim(note, "\n")

	notes = append(notes, note)

	color.HiGreen("\nadded new note!")

	return notes
}

func Delete(notes []string) []string {
	color.Magentaln("\nnotes:")
	if len(notes) == 0 {
		color.Red("\tno notes to delete")
		return notes
	}

	reader := bufio.NewReader(os.Stdin)

	for i, v := range notes {
		color.Blue("\t[%d] ", i)
		fmt.Printf("%s\n", v)
	}

	for {
		color.Magenta("\nplease enter the id of the note you want to delete: ")
		inputtedId, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputtedId = u.TrimString(inputtedId)
		id, err := strconv.Atoi(inputtedId)
		if err != nil || id >= len(notes) || id < 0 {
			continue
		}

		color.HiGreen("\ndeleting note...")
		notes = removeIndex(notes, id)
		color.HiGreen("\ndone!")

		return notes
	}
}

func Clear() ([]string, bool) {
	var (
		notes   []string
		cleared bool
	)

	color.Set(color.FgHiMagenta)
	response := yon.Prompt("\nare you sure you want to delete all your note")
	color.Unset()
	if response == yon.Yes {
		color.HiGreen("\nall notes deleted!")
		cleared = true
		return notes, cleared
	} else {
		color.Red("\ncancelled...")
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

func Check() {
	if _, err := os.Stat(note.File); os.IsNotExist(err) {
		if _, err := os.Stat(note.Directory); os.IsNotExist(err) {
			dirErr := os.Mkdir(note.Directory, 0755)
			if dirErr != nil {
				log.Fatal(dirErr)
			}
		}

		var notes []string

		Overwrite(notes)
	}
}

func Read() []string {
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

func removeIndex(notes []string, index int) []string {
	return append(notes[:index], notes[index+1:]...)
}
