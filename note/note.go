package note

import (
	"fmt"
	"log"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	home      = getHomedir()
	dataDir   = "gemdotData"
	fileName  = "notes.json"
	Directory = fmt.Sprintf("%s/%s/", home, dataDir)
	File      = fmt.Sprintf("%s%s", Directory, fileName)
)

func getHomedir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}
