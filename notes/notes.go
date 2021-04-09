package notes

import (
	"log"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	home      = getHomedir()
	Directory = home + "/gemdotData/"
	File      = home + "/gemdotData/notes.json"
)

func getHomedir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}
